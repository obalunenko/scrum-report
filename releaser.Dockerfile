FROM alpine:3.14.0
RUN apk add -U --no-cache ca-certificates

RUN mkdir -p /data/input && \
    mkdir -p /data/result && \
    mkdir -p /data/archive

COPY scrum-report /

ENTRYPOINT ["/scrum-report"]