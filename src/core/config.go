// Defines and initiates a Configuration struct
package core

import "os"

type settings struct {
	Api      ApiConfig
	Server   ServerConfig
	Database DatabaseConfig
}

type ApiConfig struct {
	Name    string
	Version string
	Mode    string
}

type ServerConfig struct {
	Host string
	Port string
}

type DatabaseConfig struct {
	Dialect string
	Uri     string
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var Config = settings{
	Api: ApiConfig{
		Name:    getEnv("API_NAME", "go-api"),
		Version: getEnv("API_VERSION", "v1"),
		Mode:    getEnv("GIN_MODE", "debug"),
	},
	Server: ServerConfig{
		Host: getEnv("HOST", "0.0.0.0"),
		Port: getEnv("PORT", "8000"),
	},
	Database: DatabaseConfig{
		Dialect: getEnv("DATABASE_DIALECT", "sqlite"),
		Uri:     getEnv("DATABASE_URI", "main.db"),
	},
}
