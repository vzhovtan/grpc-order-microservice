package config

import (
	"log"
	"os"
	"strconv"
)

func GetEnv() string {
	return getEnvValue("ENV")
}

func GetDataSrc() string {
	return getEnvValue("DATA_SOURCE")
}
func GetPaymentServiceUrl() string {
	return getEnvValue("PAYMENT_SERVICE_URL")
}

func GetPort() int {
	portStr := getEnvValue("APP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("port is invalid %s", portStr)
	}
	return port
}

func getEnvValue(key string) string {
	if os.Getenv(key) == "" {
		log.Fatalf("env variable is missing %s", key)
	}
	return os.Getenv(key)
}
