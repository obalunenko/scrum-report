package main

import (
	"log"

	"github.com/oleg-balunenko/scrum-report/config"
	"github.com/oleg-balunenko/scrum-report/logger"
	"github.com/oleg-balunenko/scrum-report/reporter"
)

func main() {
	printVersion()

	cfg := config.Load()

	logger.SetUp(cfg)
	r := reporter.New(cfg)

	log.Fatal(r.Run())

}
