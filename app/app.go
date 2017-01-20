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

type App struct {
	appName string
	hc      HttpConfig

	db        db.DB
	dbc       *pg.DB
	echo      *echo.Echo
	warn, log *log.Logger

	statTotalHits *prometheus.CounterVec
}

func (a *App) registerStat() {
	a.statTotalHits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: a.appName,
		Subsystem: "sample",
		Name:      "hits_count",
		Help:      "Total sample hits.",
	}, nil)

	prometheus.MustRegister(a.statTotalHits)
	a.echo.Any("/metrics", echo.WrapHandler(promhttp.Handler()))
}

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

func (a *App) Printf(format string, v ...interface{}) {
	if a.log != nil {
		a.log.Output(2, fmt.Sprintf(format, v...))
	}
}

func (a *App) Errorf(format string, v ...interface{}) {
	if a.warn != nil {
		a.warn.Output(2, fmt.Sprintf(format, v...))
	}
}

func (a *App) Run() error {
	// go a.process1()
	a.registerStat()

	return a.runHttpHandler(a.hc.Host, a.hc.Port)
}
