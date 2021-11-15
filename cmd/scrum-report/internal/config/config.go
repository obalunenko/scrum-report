// Package config provide application configuration.
package config

import (
	"flag"
)

var (
	// Log related configs.
	logLevel              = flag.String("log_level", "INFO", "set log level of application")
	logFormat             = flag.String("log_format", "text", "Format of logs (supported values: text, json")
	logSentryDSN          = flag.String("log_sentry_dsn", "", "Sentry DSN")
	logSentryTraceEnabled = flag.Bool("log_sentry_trace_enable", false,
		"Enables sending stacktrace to sentry")
	logSentryTraceLevel = flag.String("log_sentry_trace_level", "PANIC",
		"The level at which to start capturing stacktraces")

	appName = flag.String("app_name", "scrum-report", "app service name")
	appPort = flag.String("app_port", "8080", "app port")
)

// ensureFlags panics if env is checked before flags are parsed.
// Ok to panic since this should be caught in dev or staging.
func ensureFlags() {
	if !flag.Parsed() {
		panic("flags not parsed yet")
	}
}

// LogLevel config.
func LogLevel() string {
	ensureFlags()
	return *logLevel
}

// LogSentryDSN config.
func LogSentryDSN() string {
	ensureFlags()
	return *logSentryDSN
}

// LogSentryEnabled config.
func LogSentryEnabled() bool {
	ensureFlags()
	return LogSentryDSN() != ""
}

// LogSentryTraceEnabled config.
func LogSentryTraceEnabled() bool {
	ensureFlags()
	return *logSentryTraceEnabled
}

// LogSentryTraceLevel config.
func LogSentryTraceLevel() string {
	ensureFlags()
	return *logSentryTraceLevel
}

// LogFormat config.
func LogFormat() string {
	ensureFlags()
	return *logFormat
}

// AppPort config.
func AppPort() string {
	ensureFlags()
	return *appPort
}

// AppName config.
func AppName() string {
	ensureFlags()
	return *appName
}

// Load loads application configuration.
func Load() {
	flag.Parse()
}
