package app

import (
	"fmt"
	_ "net/http/pprof"
)

func (a *App) runHttpHandler(host string, port int) error {
	listenAddress := fmt.Sprintf("%s:%d", host, port)
	a.Printf("starting http listener at http://%s\n", listenAddress)

	return a.echo.Start(listenAddress)
}
