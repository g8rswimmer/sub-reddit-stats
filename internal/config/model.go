package config

// Settings are the configuration settings that are used to set up variables in
// the services.
type Settings struct {
	Database *Database `json:"database"`
	Reddit   *Reddit   `json:"reddit"`
	Server   *Server   `json:"server"`
}

// Database is the the configuration of the database.  SQLite is used,
// is the data source will either be in memory or a file.
type Database struct {
	DataSource string `json:"data_source"`
}

// Reddit is the configuration for the reddit API client and
// oauth.
type Reddit struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	BaseURL      string `json:"base_url"`
	Subreddit    string `json:"subreddit"`
}

// Server is the configuration for the ports for gRPC and HTTP servers.
type Server struct {
	GRPCPort int `json:"grpc_port"`
	HTTPPort int `json:"http_port"`
}
