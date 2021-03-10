package internal

import (
	"os"
)

// GetEnv will try to get value from os.LookupEnv and otherwise return fallback string
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
