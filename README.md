# stubsrv

[![goreport](https://goreportcard.com/badge/github.com/sergeyfast/stubsrv)](https://goreportcard.com/report/github.com/sergeyfast/stubsrv)

`stubsrv` is a skeleton application in Go for building simple services.

# Build on top of

* [labstack/echo](https://github.com/labstack/echo) high performance, minimalist Go web framework
* [BurntSushi/toml](https://github.com/BurntSushi/toml) TOML config parser
* [go-pg/pg](https://github.com/go-pg/pg) PostgreSQL ORM for with focus on PostgreSQL features and performance
* [prometheus/client_golang](https://github.com/prometheus/client_golang) [Prometheus](https://prometheus.io) instrumentation library for Go applications 

# Tooling

* Testing: [smartystreets/goconvey](https://github.com/smartystreets/goconvey)
* Vendoring: [rancher/trash](https://github.com/rancher/trash)
* Static analyzer: [staticcheck](https://honnef.co/go/staticcheck/cmd/staticcheck)

# Overview

This sample app was created as proof-of-concept template for new `services`. It uses proven packages and following 
[Twelve-Factor App](https://12factor.net) principles.

## Makefile

* `make tools` installs goconvey, trash and staticcheck into your `GOPATH`
* `make fix` runs `go get .`
* `make vet` runs `go vet` and `staticcheck`
* `make fmt` runs `gofmt`
* `make deps` runs `trash` for vendor.conf (don't forget to specify correct versions of libraries)
* `make build` or `make rebuild` builds or rebuilds your app
* `make run` runs app in verbose mode with default `config.cfg`
* `make test` or `make test-short` runs `go test` (with or without -test.short flag)
* `make convey` runs `goconvey` for testing and code coverage.

## First run

* Set correct database credentials in config.cfg
* run `make run`
* open browser at http://localhost:8080/sample-url and you will see bloated tables in your database :).
 Also you can check:
    * http://localhost:8080/debug/pprof/ standard /debug/pprof handler
    * http://localhost:8080/metrics url for Prometheus collector

# What's inside?

* basic structure and code organization
* TOML configs
* standard log.Logger usage with two levels: debug & error
* working with PostgreSQL database
* working with labstack/echo
* working with metrics via Prometheus (labstack/echo metrics as middleware)
* working with tests via goconvey
* vendoring example via trash

# License

This project is released under the MIT license.