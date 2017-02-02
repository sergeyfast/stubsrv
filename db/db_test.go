package db_test

import (
	"github.com/BurntSushi/toml"
	"github.com/sergeyfast/stubsrv/db"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/pg.v5"
	"log"
	"os"
	"testing"
)

var cfg struct {
	Database *pg.Options
}

func TestMain(m *testing.M) {
	if _, err := toml.DecodeFile("../config.cfg", &cfg); err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func TestDB(t *testing.T) {
	Convey("Create database instance", t, func() {
		db := db.New(pg.Connect(cfg.Database))
		Convey("Run 'select version()' query", func() {
			v, err := db.Version()
			So(err, ShouldBeNil)
			So(v, ShouldContainSubstring, "PostgreSQL")
		})

		Convey("Run bloated tables query", func() {
			list, err := db.BloatedTables()
			So(err, ShouldBeNil)
			So(list, ShouldNotBeEmpty)
			So(list[0], ShouldNotBeNil)
			So(list[0].WastedBytes, ShouldBeGreaterThan, 0)
		})

	})
}
