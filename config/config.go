package config

import (
	"fmt"
)

type IConfig interface {
	GetPort() string
	GetRedisConfig() RedisConfig
	GetCors() CorsConfig
	GetServerConfig() ServerConfig
}

type ServerConfig struct {
	IdleTimeoutMult int
	ReadTimeoutMult int
	WriteTimeout    int
	JwtSecret       []byte
}

type CorsConfig struct {
	TrustedOrigins []string
}

type RedisConfig struct {
	Addr     string
	Password string
	Username string
}

// GetConfig retrieves configuration based on the provided environment.
func GetConfig(envPtr string) (IConfig, error) {
	// In this example, we'll create sample configurations for different environments.
	// You can adjust these values based on your actual needs.
	configs := map[string]IConfig{
		"development": DevelopmentConfig{},
		"production":  ProductionConfig{},
	}

	config, ok := configs[envPtr]
	if !ok {
		return nil, fmt.Errorf("environment '%s' not found", envPtr)
	}

	return config, nil
}

// DevelopmentConfig is an example configuration for the development environment.
type DevelopmentConfig struct{}

func (dc DevelopmentConfig) GetPort() string {
	return "8080"
}

func (dc DevelopmentConfig) GetRedisConfig() RedisConfig {
	return RedisConfig{
		Addr:     "localhost:6379",
		Password: "",
		Username: "",
	}
}

func (dc DevelopmentConfig) GetCors() CorsConfig {
	return CorsConfig{
		TrustedOrigins: []string{"http://localhost:3000"},
	}
}

func (dc DevelopmentConfig) GetServerConfig() ServerConfig {
	return ServerConfig{
		IdleTimeoutMult: 1,
		ReadTimeoutMult: 2,
		WriteTimeout:    30,
		JwtSecret:       []byte("development-secret"),
	}
}

// ProductionConfig is an example configuration for the production environment.
type ProductionConfig struct{}

func (pc ProductionConfig) GetPort() string {
	return "80"
}

func (pc ProductionConfig) GetRedisConfig() RedisConfig {
	return RedisConfig{
		Addr:     "production-redis:6379",
		Password: "secret",
		Username: "user",
	}
}

func (pc ProductionConfig) GetCors() CorsConfig {
	return CorsConfig{
		TrustedOrigins: []string{"https://example.com"},
	}
}

func (pc ProductionConfig) GetServerConfig() ServerConfig {
	return ServerConfig{
		IdleTimeoutMult: 2,
		ReadTimeoutMult: 5,
		WriteTimeout:    60,
		JwtSecret:       []byte("production-secret"),
	}
}
