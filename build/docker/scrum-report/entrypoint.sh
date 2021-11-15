#!/bin/sh

set -e

echo "current user $(whoami)"

./scrum-report\
  --log_level="${SCRUM_REPORT_LOG_LEVEL}" \
  --log_format="${SCRUM_REPORT_LOG_FORMAT}" \
  --log_sentry_dsn="${SCRUM_REPORT_LOG_SENTRY_DSN}" \
  --log_sentry_trace_enable="${SCRUM_REPORT_LOG_SENTRY_TRACE_ENABLE}" \
  --log_sentry_trace_level="${SCRUM_REPORT_LOG_SENTRY_TRACE_LEVEL}" \
  --app_port="${SCRUM_REPORT_APP_PORT}" \
  --app_name="${SCRUM_REPORT_APP_NAME}" \
