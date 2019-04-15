package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/oleg-balunenko/scrum-report/config"
)

func main() {
	printVersion()

	cfg := config.Load()

	setupLogger(cfg)
	r := NewRouter()

	addr := fmt.Sprintf(":%s", cfg.Port)

	log.Fatal(http.ListenAndServe(addr, r))
}
