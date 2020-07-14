package config

import (
    "os"
)

// GetEnvVar gets a `key` from the environment variables
func GetEnvVar(key string) string {
    return os.Getenv(key)
}

func GetEnvVarDefault(key string, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    } else {
        return fallback
    }
}
