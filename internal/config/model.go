package config

type Settings struct {
	Database *Database `json:"database"`
	Reddit   *Reddit   `json:"reddit"`
	Server   *Server   `json:"server"`
}

type Database struct {
	DataSource string `json:"data_source"`
}

type Reddit struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	BaseURL      string `json:"base_url"`
	Subreddit    string `json:"subreddit"`
}

type Server struct {
	GRPCPort int `json:"grpc_port"`
	HTTPPort int `json:"http_port"`
}
