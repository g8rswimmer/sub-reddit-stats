package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/g8rswimmer/sub-reddit-stats/internal/config"
	"github.com/g8rswimmer/sub-reddit-stats/internal/datastore"
	"github.com/g8rswimmer/sub-reddit-stats/internal/manager"
	"github.com/g8rswimmer/sub-reddit-stats/internal/proto/redditv1"
	"github.com/g8rswimmer/sub-reddit-stats/internal/service"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/mvrilo/go-redoc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	cfgFileName := flag.String("config", "", "config file for migration")
	flag.Parse()

	cfg, err := config.SettingFromFile(*cfgFileName)
	if err != nil {
		slog.Error("unable to load configuration settings", "error", err.Error())
		panic(err)
	}

	slog.Info("starting server init...")
	db, err := datastore.Open(cfg.Database.DataSource)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fetcher := &datastore.Listing{
		DB: db,
	}

	mngr := &manager.Reddit{
		Fetcher: fetcher,
	}

	srv := &service.Reddit{
		Manager: mngr,
	}

	gServer := gRPCServer()
	redditv1.RegisterRedditServiceServer(gServer, srv)

	srvr, err := setUpHTTPServer(context.Background(), cfg.Server.HTTPPort, cfg.Server.GRPCPort)
	if err != nil {
		panic(err)
	}

	go func() {
		slog.Info("starting gRPC server...")
		if err := gRPCRun(gServer, cfg.Server.GRPCPort); !errors.Is(err, grpc.ErrServerStopped) && err != nil {
			panic(err)
		}
	}()

	go func() {
		slog.Info("starting HTTP server...")
		if err := srvr.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) && err != nil {
			panic(err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	slog.Info("stopping HTTP server...")
	srvr.Shutdown(context.Background())

	slog.Info("stoping gRPC server...")
	gServer.GracefulStop()
}

func gRPCRun(s *grpc.Server, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("unabled to listen to port %d: %w", port, err)
	}

	return s.Serve(lis)
}

func gRPCServer() *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionAge:      time.Second,
			MaxConnectionAgeGrace: 20 * time.Second,
		}),
	}

	return grpc.NewServer(opts...)
}

func gatewayMux() *runtime.ServeMux {
	return runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   true,
				EmitUnpopulated: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
	)
}

func setUpRouter(ctx context.Context, port int) (*mux.Router, error) {
	router := mux.NewRouter().StrictSlash(true)
	rmux := gatewayMux()
	router.PathPrefix("/").Handler(rmux)

	swagger := redoc.Redoc{
		Title:       "Subreddit",
		Description: "Subreddit Stats",
		SpecFile:    "./swagger/protos/reddit/reddit.swagger.json",
		SpecPath:    "/docs/reddit.json",
	}
	rmux.HandlePath(http.MethodGet, "/docs", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		swagger.Handler().ServeHTTP(w, r)
	})
	rmux.HandlePath(http.MethodGet, "/docs/reddit.json", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		swagger.Handler().ServeHTTP(w, r)
	})

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	if err := redditv1.RegisterRedditServiceHandlerFromEndpoint(ctx, rmux, fmt.Sprintf("localhost:%d", port), opts); err != nil {
		return nil, fmt.Errorf("reddit service handler from endpoint err: %w", err)
	}

	return router, nil
}

func setUpHTTPServer(ctx context.Context, httpPort, grpcPort int) (*http.Server, error) {
	router, err := setUpRouter(ctx, grpcPort)
	if err != nil {
		return nil, err
	}

	return &http.Server{
		Addr:              fmt.Sprintf(":%d", httpPort),
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       20 * time.Second,
		WriteTimeout:      20 * time.Second,
		Handler:           router,
	}, nil
}
