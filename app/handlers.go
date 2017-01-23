package app

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	_ "net/http/pprof"
)

// runHttpHandler is a function that starts http listener using labstack/echo.
func (a *App) runHttpHandler(host string, port int) error {
	listenAddress := fmt.Sprintf("%s:%d", host, port)
	a.Printf("starting http listener at http://%s\n", listenAddress)

	return a.echo.Start(listenAddress)
}

// registerHttpHandlers is a function that adds handlers to a.echo instance.
func (a *App) registerHttpHandlers() {
	a.echo.GET("/sample-url", func(c echo.Context) error {
		if list, err := a.db.BloatedTables(); err != nil {
			return c.HTML(http.StatusInternalServerError, err.Error())
		} else {
			a.statTotalHits.WithLabelValues().Inc() // sample call
			return c.JSON(http.StatusOK, list)
		}
	})
}
