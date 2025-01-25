package envs

import (
	"os"
	"strconv"
)

func GetEnv(key, defaultValue string) string {
	// This will by-default provide the string value
	val, found := os.LookupEnv(key)
	if !found {
		return defaultValue
	}
	return val
}

func GetEnvAsInt(key string, defaultValue int) int {
	valStr := GetEnv(key, "")
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}
	return defaultValue
}
