package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/g8rswimmer/sub-reddit-stats/internal/datastore"
	"github.com/g8rswimmer/sub-reddit-stats/internal/manager"
	"github.com/g8rswimmer/sub-reddit-stats/internal/proto/redditv1"
	"github.com/g8rswimmer/sub-reddit-stats/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func main() {
	slog.Info("starting server init...")
	db, err := datastore.Open("./db/sqlite-database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fetcher := &datastore.Presister{
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

	go func() {
		slog.Info("starting gRPC server...")
		if err := gRPCRun(gServer, 5050); !errors.Is(err, grpc.ErrServerStopped) && err != nil {
			panic(err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	gServer.GracefulStop()
	slog.Info("stoping gRPC server...")
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
