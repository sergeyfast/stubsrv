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

	statLogEvents *prometheus.CounterVec
}

// registerMetrics is a function that initializes a.stat* variables and adds /metrics endpoint to echo.
func (a *App) registerMetrics() {
	a.statLogEvents = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: a.appName,
		Subsystem: "log",
		Name:      "events_count",
		Help:      "Log events distributions.",
	}, []string{"type"})

	prometheus.MustRegister(a.statLogEvents)

	a.echo.Use(HTTPMetrics(a.appName))
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
	if a.statLogEvents != nil {
		a.statLogEvents.WithLabelValues("debug").Inc()
	}

	if a.log != nil {
		a.log.Output(2, fmt.Sprintf(format, v...))
	}
}

// Errorf prints message to Stderr (app.wan variable).
func (a *App) Errorf(format string, v ...interface{}) {
	if a.statLogEvents != nil {
		a.statLogEvents.WithLabelValues("error").Inc()
	}

	if a.warn != nil {
		a.warn.Output(2, fmt.Sprintf(format, v...))
	}
}

// Run is a function that runs application.
func (a *App) Run() error {
	// go a.process1()
	a.registerMetrics()
	a.registerHTTPHandlers()
	a.registerDebugHandlers()

	return a.runHTTPServer(a.hc.Host, a.hc.Port)
}
