package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configuration *Config

type DbConfig struct {
	Host          string
	Port          int
	Name          string
	User          string
	Password      string
	SslEnableMode bool
}

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
	DB           DbConfig
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

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("DB Host is required")
		return
	}
	dbPrt := os.Getenv("DB_PORT")
	if dbPrt == "" {
		fmt.Println("DB Port is required")
		return
	}
	dbPort, err := strconv.Atoi(dbPrt)
	if err != nil {
		os.Exit(1)
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("DB Name is required")
		return
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("DB User is required")
		return
	}
	dbPass := os.Getenv("DB_PASSWORD")
	if dbPass == "" {
		fmt.Println("DB Password is required")
		return
	}
	dbPassword := dbPass
    sslMode := os.Getenv("DB_SSL_ENABLE_MODE")
	dbSslMode := sslMode == "true"

	dbConfig := &DbConfig{
		Host: dbHost,
		Port: dbPort,
		Name: dbName,
		User: dbUser,
		Password: dbPassword,
		SslEnableMode: dbSslMode,
	}

	configuration = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     httpPort,
		JwtSecretKey: jwtSecretKey,
		DB: *dbConfig,
	}

}

func GetConfig() *Config {
	if configuration == nil {
		loadConfig()
	}
	return configuration
}
