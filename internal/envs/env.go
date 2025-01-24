package envs

import "os"

func GetEnv(key, defaultValue string) string {
	val, found := os.LookupEnv(key)
	if !found {
		return defaultValue
	}
	return val
}
