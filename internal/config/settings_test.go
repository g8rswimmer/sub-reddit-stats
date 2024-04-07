package config

import (
	"reflect"
	"testing"
)

func TestSettingFromFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    *Settings
		wantErr bool
	}{
		{
			name: "settings",
			args: args{
				filename: "config-test.json",
			},
			want: &Settings{
				Database: &Database{
					DataSource: "./db/sqlite-database.db",
				},
				Reddit: &Reddit{
					ClientID:     "your client ID",
					ClientSecret: "your client secret",
					BaseURL:      "https://oauth.reddit.com",
					Subreddit:    "funny",
				},
				Server: &Server{
					GRPCPort: 5050,
					HTTPPort: 8080,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SettingFromFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("SettingFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SettingFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
