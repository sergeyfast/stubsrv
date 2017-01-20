package db

import (
	"gopkg.in/pg.v5"
)

type DB struct {
	*pg.DB
}

func New(db *pg.DB) DB {
	return DB{db}
}

func (db DB) Version() (string, error) {
	var v string
	if _, err := db.QueryOne(pg.Scan(&v), "select version()"); err != nil {
		return "", err
	}

	return v, nil
}
