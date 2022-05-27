package configuration

import "os"

const (
	GoEnvironment          = "GO_ENVIRONMENT"
	EnvironmentProduction  = "production"
	EnvironmentDevelopment = "develop"

	Port = "PORT"
)

func getEnvironment() string {
	return os.Getenv(GoEnvironment)
}

func getWd() string {
	wd, _ := os.Getwd()
	return wd
}

func getPort() string {
	port := os.Getenv(Port)
	if port == "" {
		port = ":8080"
	}
	return port
}
