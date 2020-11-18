package main

import (
	"log"

	"github.com/obalunenko/scrum-report/internal/config"
	"github.com/obalunenko/scrum-report/internal/logger"
	"github.com/obalunenko/scrum-report/internal/reporter"
)

func main() {
	printVersion()

	cfg := config.Load()

	logger.SetUp(cfg)

	r := reporter.New(cfg)

	log.Fatal(r.Run())
}
