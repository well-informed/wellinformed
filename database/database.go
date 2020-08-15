package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed"
)

type DB struct {
	*sqlx.DB //embeds the sql db methods on the DB struct
}

/*NewDB Creates a new handle on the database
and creates necessary tables if they do not already exist*/
func NewDB(conf wellinformed.Config) DB {
	format := "host=%v  dbname=%v user=%v password=%v sslmode=disable"
	connStr := fmt.Sprintf(format, conf.DBHost, conf.DBName, conf.DBUser, conf.DBPassword)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal("could not connect to database. err: ", err)
	}
	createTables(db, tables)

	return DB{db}
}

/*Creates all necessary tables, either returns successfully,
or exits the program with call to log.Fatal()*/
func createTables(db *sqlx.DB, tables []table) {
	for _, table := range tables {
		createTable(db, table.name, table.sql)
	}
}

func createTable(db *sqlx.DB, name string, stmt string) {
	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatalf("error creating table %v. err: %v", name, err)
	}
}
