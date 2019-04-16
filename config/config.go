package config

import (
	"flag"
)

// Config stores configuration for service
type Config struct {
	LogLevel    string
	Port        string
	Host        string
	OpenBrowser bool
}

// Load configuration from flags
func Load() *Config {
	var c Config
	flag.StringVar(&c.Host, "host_address", "localhost", "address of host")
	flag.StringVar(&c.Port, "listen_port", "8080", "listen port")
	flag.StringVar(&c.LogLevel, "log_level", "INFO", "log level")
	flag.BoolVar(&c.OpenBrowser, "open_browser", false, "open browser after start on index page")

	flag.Parse()

	return &c
}
