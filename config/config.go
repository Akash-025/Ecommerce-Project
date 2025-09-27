package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configuration *Config

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
}

func loadConfig() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load the env variable", err)
		os.Exit(1)
	}
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
		return
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("ServiceName is required")
		return
	}
	hp := os.Getenv("HTTP_PORT")
	if hp == "" {
		fmt.Println("HttpPort is required")
		return
	}
	httpPort, err := strconv.Atoi(hp)
	if err != nil {
		fmt.Println("HttpPort must be number")
		return
	}
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	configuration = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     httpPort,
		JwtSecretKey: jwtSecretKey,
	}

}

func GetConfig() *Config {
	if configuration == nil {
		loadConfig()
	}
	return configuration
}
