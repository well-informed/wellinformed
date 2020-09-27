package database

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed"
	"github.com/well-informed/wellinformed/database/migrations"
)

type DB struct {
	*sqlx.DB //embeds the sql db methods on the DB struct
}

/*NewDB Creates a new handle on the database
and creates necessary tables if they do not already exist*/
func NewDB(conf wellinformed.Config) DB {
	format := "postgres://%v:%v@%v:5432/%v?sslmode=disable"
	connStr := fmt.Sprintf(format, conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBName)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal("could not connect to database. err: ", err)
	}
	db.DB.SetMaxOpenConns(conf.DBMaxOpenConnections)
	db.DB.SetMaxIdleConns(conf.DBMaxIdleConnections)
	MigrateSchema(connStr)
	return DB{db}
}

//Handles schema migration by reading binary packed sql files from migrations/bindata.go
func MigrateSchema(dbURL string) {
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
	if err == migrate.ErrNoChange {
		log.Warn("bypassing migration: ", err)
	}
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("could not run migration: ", err)
	}
}
