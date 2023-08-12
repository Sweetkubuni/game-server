package config

import (
	"github.com/caarlos0/env/v8"
)

var Instance Config

// Config is the configuration of the application
type Config struct {
	Server ServerConfig
	API    APIConfig
}

// ServerConfig is the configuration for the server
type ServerConfig struct {
	Host         string `env:"HOST" envDefault:"localhost"`
	Port         string `env:"PORT" envDefault:"8080"`
	IsProduction bool   `env:"IS_PRODUCTION" envDefault:"false"`
	JWTSecret    string `env:"JWT_SECRET" envDefault:"supersecretkey"`
}

// APIConfig is the configuration  for API
type APIConfig struct {
	RoomNameLength        int `env:"ROOMNAME_LENGTH,notEmpty" envDefault:"20"`
	MinRoomNameLength     int `env:"MIN_ROOMNAME_LENGTH,notEmpty" envDefault:"2"`
	RoomDescriptionLength int `env:"ROOM_DESCRIPTION_LENGTH,notEmpty" envDefault:"200"`
}

// LoadConfig loads the configuration from environment variables
func LoadConfig() error {
	var server ServerConfig
	if err := env.Parse(&server); err != nil {
		return err
	}

	var apiConfig APIConfig
	if err := env.Parse(&apiConfig); err != nil {
		return err
	}

	cfg := &Config{
		Server: server,
		API:    apiConfig,
	}

	Instance = *cfg

	return nil
}
