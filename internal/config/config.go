package config

import (
	"os"
)

// Config stores configuration for service.
type Config struct {
	LogLevel string
	Port     string
}

// Load configuration from flags.
func Load() Config {
	var c Config

	c.Port = getStringOrDefault("SCRUM_REPORT_PORT", "8080")
	c.LogLevel = getStringOrDefault("SCRUM_REPORT_LOG_LEVEL", "INFO")

	return c
}

func getStringOrDefault(key string, defVal string) string {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		return defVal
	}

	return val
}
