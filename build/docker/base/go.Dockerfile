FROM golang:1.19.3-alpine3.16
LABEL maintainer="oleg.balunenko@gmail.com"
LABEL org.opencontainers.image.source="https://github.com/obalunenko/scrum-report"
LABEL stage="base"

ARG PROJECT_URL=github.com/obalunenko/scrum-report
RUN mkdir -p "${GOPATH}/src/${PROJECT_URL}/base-tools"

WORKDIR "${GOPATH}/src/${PROJECT_URL}/base-tools"

RUN apk update && \
    apk add --no-cache \
        "git" \
        "make" \
        "gcc" \
        "bash" \
        "curl" \
        "musl-dev" \
        "unzip" \
        "ca-certificates" \
        "libstdc++" \
        "binutils-gold" && \
    rm -rf /var/cache/apk/*

# Get and install glibc for alpine
ARG APK_GLIBC_VERSION=2.29-r0
ARG APK_GLIBC_FILE="glibc-${APK_GLIBC_VERSION}.apk"
ARG APK_GLIBC_BIN_FILE="glibc-bin-${APK_GLIBC_VERSION}.apk"
ARG APK_GLIBC_BASE_URL="https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${APK_GLIBC_VERSION}"
# hadolint ignore=DL3018
RUN wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub \
    && wget -nv "${APK_GLIBC_BASE_URL}/${APK_GLIBC_FILE}" \
    && apk --no-cache add "${APK_GLIBC_FILE}" \
    && wget -nv "${APK_GLIBC_BASE_URL}/${APK_GLIBC_BIN_FILE}" \
    && apk --no-cache add "${APK_GLIBC_BIN_FILE}" \
    && rm glibc-*

COPY .git .git
COPY scripts scripts
COPY tools tools

COPY Makefile Makefile

# install tools from vendor
RUN make install-tools && \
    rm -rf "${GOPATH}/src/${PROJECT_URL}/base-tools"

ENV GOBIN="${GOPATH}/bin"
ENV PATH="${PATH}":"${GOBIN}"
