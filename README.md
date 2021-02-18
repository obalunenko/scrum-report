# scrum-report


[![GO](https://img.shields.io/github/go-mod/go-version/obalunenko/scrum-report)](https://golang.org/doc/devel/release.html)
![Lint & Test & Build & Release](https://github.com/obalunenko/scrum-report/workflows/Lint%20&%20Test%20&%20Build%20&%20Release/badge.svg)
![Lint](https://github.com/obalunenko/scrum-report/workflows/Lint/badge.svg)
![Test & Build](https://github.com/obalunenko/scrum-report/workflows/Test%20&%20Build/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/obalunenko/scrum-report)](https://goreportcard.com/report/github.com/obalunenko/scrum-report)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=oleg-balunenko_scrum-report&metric=alert_status)](https://sonarcloud.io/dashboard?id=oleg-balunenko_scrum-report)
[![Latest release artifacts](https://img.shields.io/github/v/release/obalunenko/scrum-report)](https://github.com/obalunenko/scrum-report/releases/latest)
[![Docker pulls](https://img.shields.io/docker/pulls/olegbalunenko/scrum-report)](https://hub.docker.com/r/olegbalunenko/scrum-report)
[![License](https://img.shields.io/github/license/obalunenko/scrum-report)](/LICENSE)

Daily stand up meeting scrum report generator in markdown format for slack

## Template

    ```text
        *What did I do*
        •
        *What will I do*
        •
        *Impediments*
        •
    ```

## Usage of scrum-report

### Run via binary
Download archive with binary from [![Latest release artifacts](https://img.shields.io/badge/artifacts-download-blue.svg)](https://github.com/obalunenko/scrum-report/releases/latest)
unzip it and execute binary `scrum-report`
Following parameters could be configured via environment variables:

    ```bash
        SCRUM_REPORT_PORT=8080
        SCRUM_REPORT_LOG_LEVEL=INFO
    ```

Then open in browser page `localhost:8080` and start to us it

### Run via docker-compose
Download archive with docker-compose file from [![Latest release artifacts](https://img.shields.io/badge/artifacts-download-blue.svg)](https://github.com/obalunenko/scrum-report/releases/latest)
Unzip it and execute in unzipped directory:

`docker-compose up`

Then open in browser page `localhost:8080` and start to us it

### Run via docker
Pull latest image from [![docker hub]](https://hub.docker.com/r/olegbalunenko/scrum-report)

    ```bash 
        docker run -d -t -i -e SCRUM_REPORT_PORT='8080' -e SCRUM_REPORT_LOG_LEVEL='INFO' \
        -p 8080:8080 \
        --name scrum-report olegbalunenko/scrum-report
    ```
Then open in browser page `localhost:8080` and start to us it

#### Demo

![first step](.github/images/img1.png)

![first step](.github/images/img2.png)

![first step](.github/images/img3.png)
