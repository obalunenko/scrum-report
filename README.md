# scrum-report


[![GO](https://img.shields.io/github/go-mod/go-version/oleg-balunenko/scrum-report)](https://golang.org/doc/devel/release.html)
[![Build Status](https://travis-ci.org/oleg-balunenko/scrum-report.svg?branch=master)](https://travis-ci.org/oleg-balunenko/scrum-report)
[![Go Report Card](https://goreportcard.com/badge/github.com/oleg-balunenko/scrum-report)](https://goreportcard.com/report/github.com/oleg-balunenko/scrum-report)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=oleg-balunenko_scrum-report&metric=alert_status)](https://sonarcloud.io/dashboard?id=oleg-balunenko_scrum-report)
[![Latest release artifacts](https://img.shields.io/github/v/release/oleg-balunenko/scrum-report)](https://github.com/oleg-balunenko/scrum-report/releases/latest)
[![Docker pulls](https://img.shields.io/docker/pulls/olegbalunenko/scrum-report)](https://hub.docker.com/r/olegbalunenko/scrum-report)
[![License](https://img.shields.io/github/license/oleg-balunenko/scrum-report)](/LICENSE)

Daily stand up meeting scrum report generator in markdown format for slack

## Template

```text
*Yesterday*
•
*Today*
•
*Impediments*
•
```

## Usage of scrum-report

Download binary from [![Latest release artifacts](https://img.shields.io/badge/artifacts-download-blue.svg)](https://github.com/oleg-balunenko/scrum-report/releases/latest)
and execute
Following parameters could be configured via environment variables:

```bash
   SCRUM_REPORT_PORT=8080
   SCRUM_REPORT_LOG_LEVEL=INFO
```

Then open in browser page `localhost:8080` and start to us it

### Demo

![first step](./docs/img1.png)

![first step](./docs/img2.png)

![first step](./docs/img3.png)
