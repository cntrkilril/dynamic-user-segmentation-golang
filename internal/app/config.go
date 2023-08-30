package app

import (
	"github.com/go-playground/validator/v10"
	"strings"
	"time"
)

type (
	Config struct {
		Logger   Logger   `validate:"required"`
		HTTP     HTTP     `validate:"required"`
		Postgres Postgres `validate:"required"`
		Static   Static   `validate:"required"`
		Swagger  Swagger  `validate:"required"`
	}

	Logger struct {
		Level *int8 `validate:"required"`
	}

	HTTP struct {
		Host     string `validate:"required"`
		Port     string `validate:"required"`
		Protocol string `validate:"required"`
	}

	Postgres struct {
		ConnString      string        `validate:"required"`
		MaxOpenConns    int           `validate:"required"`
		ConnMaxLifetime time.Duration `validate:"required"`
		MaxIdleConns    int           `validate:"required"`
		ConnMaxIdleTime time.Duration `validate:"required"`
	}

	Static struct {
		PathToSaveHistory string `validate:"required"`
	}

	Swagger struct {
		PathToConfigFile string `validate:"required"`
		ConfigUrl        string `validate:"required"`
	}
)

func LoadConfig() (*Config, error) {

	defaultLogLevel := int8(-1)

	cfg := &Config{
		HTTP: HTTP{
			Host:     "localhost",
			Port:     "8080",
			Protocol: "http",
		},
		Logger: Logger{
			Level: &defaultLogLevel,
		},
		Postgres: Postgres{
			ConnString:      "postgresql://root:pass@127.0.0.1:5432/root?sslmode=disable",
			MaxOpenConns:    10,
			ConnMaxLifetime: 20,
			MaxIdleConns:    15,
			ConnMaxIdleTime: 30,
		},
		Static: Static{
			PathToSaveHistory: "./static/history_files",
		},
		Swagger: Swagger{
			PathToConfigFile: "./docs",
			ConfigUrl:        "http://localhost:8080/docs/swagger.yaml",
		},
	}

	err := validator.New().Struct(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func UrlHttp(config *Config) string {
	urlHttp := strings.Builder{}
	urlHttp.WriteString(config.HTTP.Protocol)
	urlHttp.WriteString("://")
	urlHttp.WriteString(config.HTTP.Host)
	urlHttp.WriteString(":")
	urlHttp.WriteString(config.HTTP.Port)
	urlHttp.WriteString("/")
	return urlHttp.String()
}
