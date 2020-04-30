package main

import (
	"log"

	"github.com/oleg-balunenko/scrum-report/internal/config"
	"github.com/oleg-balunenko/scrum-report/internal/logger"
	"github.com/oleg-balunenko/scrum-report/internal/reporter"
)

func main() {
	printVersion()

	cfg := config.Load()

	logger.SetUp(cfg)

	r := reporter.New(cfg)

	log.Fatal(r.Run())
}
