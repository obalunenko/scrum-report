FROM golang:1.16-alpine as build-container

ENV PROJECT_DIR=${GOPATH}/src/github.com/obalunenko/scrum-report

RUN apk update && \
    apk upgrade && \
    apk add --no-cache git musl-dev make gcc bash

RUN mkdir -p ${PROJECT_DIR}

COPY ./  ${PROJECT_DIR}
WORKDIR ${PROJECT_DIR}

# vet project
RUN make vet
# test project
RUN make test-docker
# compile executable
RUN make compile

RUN mkdir /app
RUN cp ./bin/scrum-report /app/scrum-report


FROM alpine:3.11.3 as deployment-container
RUN apk add -U --no-cache ca-certificates


COPY --from=build-container /app/scrum-report /scrum-report

ENTRYPOINT ["/scrum-report"]

