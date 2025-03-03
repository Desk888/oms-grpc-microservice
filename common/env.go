package common

import (
	"syscall"
)

func EnvString(key string, defaultValue string) string {
	value, ok := syscall.Getenv(key)
	if !ok {
		return defaultValue
	}
	return value
}