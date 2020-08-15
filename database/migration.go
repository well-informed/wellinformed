package database

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed/database/migrations"
)

func migrateSchema(dbURL string) {
	s := bindata.Resource(migrations.AssetNames(),
		func(name string) ([]byte, error) {
			return migrations.Asset(name)
		})
	d, err := bindata.WithInstance(s)
	if err != nil {
		log.Fatal("couldn't read migration bindata. err: ", err)
	}
	m, err := migrate.NewWithSourceInstance("go-bindata", d, dbURL)
	if err != nil {
		log.Fatal("could not establish migrations source. err: ", err)
	}
	err = m.Up()
	if err != nil {
		log.Warn("could not run migration: ", err)
	}
}
