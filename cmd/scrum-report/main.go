package main

import (
	"context"

	log "github.com/obalunenko/logger"
	"github.com/obalunenko/version"

	"github.com/obalunenko/scrum-report/cmd/scrum-report/internal/config"
	"github.com/obalunenko/scrum-report/internal/reporter"
)

func main() {
	ctx := context.Background()

	config.Load()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	log.Init(ctx, log.Params{
		Level:  config.LogLevel(),
		Format: config.LogFormat(),
		SentryParams: log.SentryParams{
			Enabled:      config.LogSentryEnabled(),
			DSN:          config.LogSentryDSN(),
			TraceEnabled: config.LogSentryTraceEnabled(),
			TraceLevel:   config.LogSentryTraceLevel(),
			Tags: map[string]string{
				"app_name":     version.GetAppName(),
				"go_version":   version.GetGoVersion(),
				"version":      version.GetVersion(),
				"build_date":   version.GetBuildDate(),
				"short_commit": version.GetShortCommit(),
			},
		},
	})

	printVersion(ctx)

	svc := reporter.New(ctx, reporter.Params{
		AppName: config.AppName(),
		Port:    config.AppPort(),
	})

	<-svc.Run()

	log.Info(ctx, "cmd: Exit")
}
