package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	Environment   string `env:"ENV" envDefault:"local"`
	DBHost        string `env:"DB_HOST" envDefault:"localhost"`
	DBPort        string `env:"DB_PORT" envDefault:"5432"`
	DBUser        string `env:"DB_USER" envDefault:"pos"`
	DBPassword    string `env:"DB_PASSWORD" envDefault:"password"`
	DBName        string `env:"DB_NAME" envDefault:"pos_db"`
	ServerPort    string `env:"SERVER_PORT" envDefault:"8000"`
	AdminEmail    string `env:"ADMIN_EMAIL" envDefault:"admin@admin.com"`
	AdminPassword string `env:"ADMIN_PASSWORD" envDefault:"password"`
	AdminName     string `env:"ADMIN_NAME" envDefault:"admin"`
}

var Environment *Config

func init() {
	// Load .env file first
	godotenv.Load()

	// Parse into struct
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	Environment = &cfg
}
