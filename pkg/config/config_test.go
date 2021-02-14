package config

import (
	"reflect"
	"testing"

	_ "github.com/sumelms/microservice-catalog/tests"
)

func TestNewConfig(t *testing.T) {
	type args struct {
		configPath string
	}

	validConfig, _ := NewConfig("config/config.yml")

	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name:    "Invalid path",
			args:    args{configPath: "config.yml"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Correct path",
			args:    args{configPath: "config/config.yml"},
			want:    validConfig,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConfig(tt.args.configPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
