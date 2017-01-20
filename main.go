package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/sergeyfast/stubsrv/app"
	"github.com/sergeyfast/stubsrv/db"
	"gopkg.in/pg.v5"
	"io/ioutil"
	"log"
	"os"
)

const appName = "stubsrv"

var (
	flVerbose    = flag.Bool("verbose", false, "enable debug output")
	flConfigPath = flag.String("config", "config.cfg", "Path to config file")
	cfg          Config
	version      string
)

// Application Config
type Config struct {
	Server   app.HttpConfig
	Database *pg.Options
}

func main() {
	flag.Parse()
	fixStdLog(*flVerbose)

	log.Printf("starting %v version=%v", appName, version)
	if _, err := toml.DecodeFile(*flConfigPath, &cfg); err != nil {
		die(err)
	}

	dbc := pg.Connect(cfg.Database)
	db := db.New(dbc)

	if v, err := db.Version(); err != nil {
		die(err)
	} else {
		log.Println(v)
	}

	a := app.New(appName, *flVerbose, cfg.Server, dbc)
	die(a.Run())
}

func fixStdLog(verbose bool) {
	if verbose {
		log.SetPrefix("D")
		log.SetOutput(os.Stdout)
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	} else {
		log.SetOutput(ioutil.Discard)
	}
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
