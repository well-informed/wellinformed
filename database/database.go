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

// func createUserHistoryTable(db *sql.DB) {
// 	stmt := `
// 	CREATE TABLE IF NOT EXISTS user_history
// 	( userID varchar,
// 		parent_feed varchar,
// 		guid varchar,
// 		trustworthiness smallint,
// 		insightfulness smallint,
// 		entertainment smallint,
// 		importance smallint,
// 		overall smallint,
// 		notes text,
// 		FOREIGN KEY (userID) REFERENCES users(userID),
// 		FOREIGN KEY (parent_feed, guid) REFERENCES main_feed(parent_feed, guid),
// 		PRIMARY KEY (userID, parent_feed, guid)
// 		)`
// 	_, err := db.Exec(stmt)
// 	if err != nil {
// 		log.Fatal("error creating history table. err: ", err)
// 	}
// }

// func createUserPrefSetTable(db *sql.DB) {
// 	stmt := `
// 	CREATE TABLE IF NOT EXISTS preference_sets
// 	( userID varchar,
// 		pref_set_name varchar,
// 		FOREIGN KEY (userID) REFERENCES users(userID),
// 		PRIMARY KEY (userID, pref_set_name)
// 	)`
// 	_, err := db.Exec(stmt)
// 	if err != nil {
// 		log.Fatal("error creating preference_sets table. err: ", err)
// 	}
// }

// func createUserSourcesTable(db *sql.DB) {
// 	stmt := `
// 	CREATE TABLE IF NOT EXISTS user_sources
// 	( userID varchar,
// 		pref_set_name varchar,
// 		source varchar,
// 		FOREIGN KEY (userID, pref_set_name) REFERENCES preference_sets(userID, pref_set_name),
// 		FOREIGN KEY (source) REFERENCES src_rss_feeds(link)
// 		PRIMARY KEY (userID, pref_set_name, source)
// 	)`
// 	_, err := db.Exec(stmt)
// 	if err != nil {
// 		log.Fatal("error creating user_sources table. err: ", err)
// 	}
// }
