package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/oleg-balunenko/scrum-report/config"
)

func setupLogger(config *config.Config) {
	log.SetOutput(os.Stdout)
	log.SetFormatter(new(log.TextFormatter))
	lvl, err := log.ParseLevel(config.LogLevel)
	if err != nil {
		log.Fatalf("Failed to parse log level. %v", err)
	}
	log.SetLevel(lvl)
}
