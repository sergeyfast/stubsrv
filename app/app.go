package app

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sergeyfast/stubsrv/db"
	"gopkg.in/pg.v5"
	"log"
	"os"
)

type HttpConfig struct {
	Host string
	Port int
}

// App is application with all dependencies.
type App struct {
	appName string
	hc      HttpConfig

	db        db.DB
	dbc       *pg.DB
	echo      *echo.Echo
	warn, log *log.Logger

	statTotalHits *prometheus.CounterVec
}

// registerStats is a function that initializes a.stat* variables and adds /metrics endpoint to echo.
func (a *App) registerStats() {
	a.statTotalHits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: a.appName,
		Subsystem: "sample",
		Name:      "hits_count",
		Help:      "Total sample hits.",
	}, nil)

	prometheus.MustRegister(a.statTotalHits)

	a.echo.Any("/metrics", echo.WrapHandler(promhttp.Handler()))
}

// New is a function that returns new application instance.
func New(appName string, verbose bool, hc HttpConfig, dbc *pg.DB) *App {
	a := &App{
		appName: appName,
		hc:      hc,
		db:      db.New(dbc),
		dbc:     dbc,
		echo:    echo.New(),
		warn:    log.New(os.Stderr, "E", log.LstdFlags|log.Lshortfile),
	}

	if verbose {
		a.log = log.New(os.Stdout, "D", log.LstdFlags|log.Lshortfile)
	}

	return a
}

// Printf prints message to Stdout (app.log variable) if a.verbose is set.
func (a *App) Printf(format string, v ...interface{}) {
	if a.log != nil {
		a.log.Output(2, fmt.Sprintf(format, v...))
	}
}

// Errorf prints message to Stderr (app.wan variable).
func (a *App) Errorf(format string, v ...interface{}) {
	if a.warn != nil {
		a.warn.Output(2, fmt.Sprintf(format, v...))
	}
}

// Run is a function that runs application.
func (a *App) Run() error {
	// go a.process1()
	a.registerStats()
	a.registerHttpHandlers()

	return a.runHttpHandler(a.hc.Host, a.hc.Port)
}
