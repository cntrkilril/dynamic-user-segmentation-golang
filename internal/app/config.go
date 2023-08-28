package app

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type (
	Config struct {
		Logger   Logger   `validate:"required"`
		HTTP     HTTP     `validate:"required"`
		Postgres Postgres `validate:"required"`
	}

	Logger struct {
		Level *int8 `validate:"required"`
	}

	HTTP struct {
		Host string `validate:"required"`
		Port string `validate:"required"`
	}

	Postgres struct {
		ConnString      string        `validate:"required"`
		MaxOpenConns    int           `validate:"required"`
		ConnMaxLifetime time.Duration `validate:"required"`
		MaxIdleConns    int           `validate:"required"`
		ConnMaxIdleTime time.Duration `validate:"required"`
	}
)

func LoadConfig() (*Config, error) {

	defaultLogLevel := int8(-1)

	cfg := &Config{
		HTTP: HTTP{
			Host: "localhost",
			Port: "8080",
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
	}

	err := validator.New().Struct(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
