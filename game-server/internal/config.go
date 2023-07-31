package internal

import (
	"bufio"
	"io/fs"
	"os"
	"strings"

	"github.com/caarlos0/env/v8"
)

var Instance Config

// Config is the configuration of the application
type Config struct {
	Server  ServerConfig
	LiveKit LiveKitConfig
	API     APIConfig
}

// ServerConfig is the configuration for the server
type ServerConfig struct {
	Host         string `env:"SERVER_HOST"`
	Port         string `env:"SERVER_PORT"`
	IsProduction bool   `env:"SERVER_IS_PRODUCTION"`
	JWTSecret    string `env:"SERVER_JWT_SECRET"`
}

// LiveKitConfig is the configuration for LiveKit
type LiveKitConfig struct {
	URL       string `env:"LIVEKIT_URL,notEmpty"`
	APIKey    string `env:"LIVEKIT_API_KEY,notEmpty"`
	APISecret string `env:"LIVEKIT_API_SECRET,notEmpty"`
}

// APIConfig is the configuration  for API
type APIConfig struct {
	RoomNameLength        int `env:"API_ROOMNAME_LENGTH,notEmpty"`
	MinRoomNameLength     int `env:"API_MIN_ROOMNAME_LENGTH,notEmpty"`
	RoomDescriptionLength int `env:"API_ROOM_DESCRIPTION_LENGTH,notEmpty"`
}

// LoadConfig loads the configuration from environment variables
func LoadConfig() error {
	checkAndSourceEnv()
	var server ServerConfig
	if err := env.Parse(&server); err != nil {
		return err
	}

	var liveKit LiveKitConfig
	if err := env.Parse(&liveKit); err != nil {
		return err
	}

	var apiConfig APIConfig
	if err := env.Parse(&apiConfig); err != nil {
		return err
	}

	cfg := &Config{
		Server:  server,
		LiveKit: liveKit,
		API:     apiConfig,
	}

	Instance = *cfg

	return nil
}

// CheckAndSource looks for .env files and injects them into the runtime
func checkAndSourceEnv() {
	// find the file
	workingDir := "./"
	fSys := os.DirFS(workingDir)

	matches, err := fs.Glob(fSys, "*.env")
	if err != nil || len(matches) == 0 {
		// cannot source the .env file
		return
	}

	for _, match := range matches {
		file, err := os.Open(match)
		if err != nil {
			continue
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if val := strings.Split(line, "="); len(val) == 2 {
				os.Setenv(val[0], val[1])

			}
		}

	}

}
