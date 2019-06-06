package config

import (
	"flag"
)

// Config stores configuration for service
type Config struct {
	LogLevel    string
	Port        string
	Host        string
	Debug       bool
	OpenBrowser bool
}

// Load configuration from flags
func Load() *Config {
	var c Config
	flag.StringVar(&c.Host, "host_address", "127.0.0.1", "address of host")
	flag.StringVar(&c.Port, "listen_port", "8080", "listen port")
	flag.StringVar(&c.LogLevel, "log_level", "INFO", "log level")
	flag.BoolVar(&c.Debug, "debug", false, `Run debug mode - use real ip machine instead localhost 
	to possible to debug with Charles`)
	flag.BoolVar(&c.OpenBrowser, "open_browser", false, "open browser after start on index page")

	flag.Parse()

	return &c
}
