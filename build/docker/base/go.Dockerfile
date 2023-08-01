FROM ghcr.io/obalunenko/go-tools:v0.14.0
LABEL maintainer="oleg.balunenko@gmail.com"
LABEL org.opencontainers.image.source="https://github.com/obalunenko/scrum-report"
LABEL stage="base"

ARG PROJECT_URL=github.com/obalunenko/scrum-report

ENV GOBIN="${GOPATH}/bin"
ENV PATH="${PATH}":"${GOBIN}"
