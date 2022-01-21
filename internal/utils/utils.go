package utils

import "os"

func GetEnvOrDefault(key, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	} else {
		return defaultValue
	}
}
