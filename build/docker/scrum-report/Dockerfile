ARG DOCKER_REPO_BASE
ARG DOCKER_GO_BASE_PROD_TAG=latest
# hadolint ignore=DL3007
FROM ${DOCKER_REPO_BASE}scrum-report-go-base-prod:${DOCKER_GO_BASE_PROD_TAG} AS build-container
LABEL maintainer="oleg.balunenko@gmail.com"
LABEL org.opencontainers.image.source="https://github.com/obalunenko/scrum-report"
LABEL stage="prod"

ENV PROJECT_DIR="${GOPATH}/src/github.com/obalunenko/scrum-report"

RUN mkdir -p "${PROJECT_DIR}"

WORKDIR "${PROJECT_DIR}"

COPY .git .git
COPY cmd cmd
COPY internal internal
COPY scripts scripts
COPY vendor vendor
COPY go.mod go.mod
COPY go.sum go.sum
COPY Makefile Makefile

# compile executable
RUN make build && \
    mkdir -p /app && \
    cp ./bin/scrum-report /app/scrum-report

COPY ./build/docker/scrum-report/entrypoint.sh /app/entrypoint.sh


FROM alpine:3.17.0 AS waiter
LABEL maintainer="oleg.balunenko@gmail.com"
LABEL org.opencontainers.image.source="https://github.com/obalunenko/scrum-report"
LABEL stage="prod"

## Add the wait script to the image
ARG WAIT_VERSION=2.9.0
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/${WAIT_VERSION}/wait /wait
RUN chmod +x /wait

FROM alpine:3.17.0 AS deployment-container
LABEL maintainer="oleg.balunenko@gmail.com"
LABEL org.opencontainers.image.source="https://github.com/obalunenko/scrum-report"
LABEL stage="prod"

# Configure least privilege user
ARG UID=1000
ARG GID=1000
RUN addgroup -S scrumreport -g ${UID} && \
    adduser -S scrumreport -u ${GID} -G scrumreport -h /home/scrumreport -s /bin/sh -D scrumreport

WORKDIR /

ARG APK_CA_CERTIFICATES_VERSION=20220614-r4

RUN apk update && \
    apk add --no-cache \
    "ca-certificates=${APK_CA_CERTIFICATES_VERSION}" && \
    rm -rf /var/cache/apk/*

## Add the wait script to the image
COPY --from=waiter /wait /wait

COPY --from=build-container /app/ /

ENTRYPOINT ["sh", "-c", "/wait && /entrypoint.sh"]

USER scrumreport
