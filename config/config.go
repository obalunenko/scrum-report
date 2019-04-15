package config

import (
	"flag"
)

// Config stores configuration for service
type Config struct {
	LogLevel string `config:"SCRUM_REPORT_LOGLEVEL"`
	Port     string `config:"SCRUM_REPORT_PORT"`
}

// Load configuration from flags
func Load() *Config {
	var c Config
	flag.StringVar(&c.Port, "listen_port", "8080", "listen port")
	flag.StringVar(&c.LogLevel, "log_level", "DEBUG", "log level")

	flag.Parse()

	return &c
}
