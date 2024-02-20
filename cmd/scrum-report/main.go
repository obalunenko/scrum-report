package main

import (
	"context"
	"os"

	log "github.com/obalunenko/logger"

	"github.com/obalunenko/scrum-report/cmd/scrum-report/internal/config"
	"github.com/obalunenko/scrum-report/internal/reporter"
)

func main() {
	ctx := context.Background()

	config.Load()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	log.Init(ctx, log.Params{
		Writer:     os.Stdout,
		Level:      config.LogLevel(),
		Format:     config.LogFormat(),
		WithSource: false,
	})

	printVersion(ctx)

	svc := reporter.New(ctx, reporter.Params{
		AppName: config.AppName(),
		Port:    config.AppPort(),
	})

	<-svc.Run()

	log.Info(ctx, "cmd: Exit")
}
