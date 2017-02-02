package app

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	_ "net/http/pprof"
)

// runHTTPServer is a function that starts http listener using labstack/echo.
func (a *App) runHTTPServer(host string, port int) error {
	listenAddress := fmt.Sprintf("%s:%d", host, port)
	a.Printf("starting http listener at http://%s\n", listenAddress)

	return a.echo.Start(listenAddress)
}

// registerDebugHandlers adds /debug/pprof handlers into a.echo instance.
func (a *App) registerDebugHandlers() {
	dbg := a.echo.Group("/debug")

	// add pprof integration
	dbg.Any("/pprof/*", func(c echo.Context) error {
		if h, p := http.DefaultServeMux.Handler(c.Request()); p != "" {
			h.ServeHTTP(c.Response(), c.Request())
			return nil
		}
		return echo.NewHTTPError(http.StatusNotFound)
	})
}

// registerHTTPHandlers is a function that adds handlers to a.echo instance.
func (a *App) registerHTTPHandlers() {
	a.echo.GET("/sample-url", func(c echo.Context) error {
		list, err := a.db.BloatedTables()
		if err != nil {
			return c.HTML(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, list)
	})
}
