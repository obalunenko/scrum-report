# scrum-report

[![Build Status](https://travis-ci.org/oleg-balunenko/scrum-report.svg?branch=master)](https://travis-ci.org/oleg-balunenko/scrum-report)
[![Go Report Card](https://goreportcard.com/badge/github.com/oleg-balunenko/scrum-report)](https://goreportcard.com/report/github.com/oleg-balunenko/scrum-report)
[![Latest release artifacts](https://img.shields.io/badge/artifacts-download-blue.svg)](https://github.com/oleg-balunenko/scrum-report/releases/latest)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=oleg-balunenko_scrum-report&metric=alert_status)](https://sonarcloud.io/dashboard?id=oleg-balunenko_scrum-report)

Daily stand up meeting scrum report generator in markdown format for slack

## Template:

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
Following parameters could be configured:

```text
   -host_address string
          address of host (default "localhost")
    -listen_port string
          listen port (default "8080")
    -log_level string
          log level (default "INFO")
    -open_browser
          open browser after start on index page

```

Then open in browser page `localhost:8080` and start to us it

### Demo

![first step](./docs/img1.png)

![first step](./docs/img2.png)

![first step](./docs/img3.png)
