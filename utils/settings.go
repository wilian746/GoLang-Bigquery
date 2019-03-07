package utils

import (
	"log"
	"os"
)

func GetEnvOrDefault(env string, defaultValue string) string {
	var value string
	value = os.Getenv(env)
	if value == "" {
		value = defaultValue
	}
	return value
}

func CheckEnvExist(environmentName string) {
	proj := os.Getenv(environmentName)
	if proj == "" {
		log.Fatalf("%s environment variable must be set.", environmentName)
	}
}
