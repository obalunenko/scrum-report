FROM alpine:3.17.0
LABEL maintainer="oleg.balunenko@gmail.com"
LABEL org.opencontainers.image.source="https://github.com/obalunenko/scrum-report"
LABEL stage="release"

# Configure least privilege user
ARG UID=1000
ARG GID=1000
RUN addgroup -S scrumreport -g ${UID} && \
    adduser -S scrumreport -u ${GID} -G scrumreport -h /home/scrumreport -s /bin/sh -D scrumreport

WORKDIR /

RUN apk update && \
    apk add --no-cache \
        "ca-certificates" && \
    rm -rf /var/cache/apk/*

COPY scrum-report /
COPY build/docker/scrum-report/entrypoint.sh /

ENTRYPOINT ["/entrypoint.sh"]

USER scrumreport