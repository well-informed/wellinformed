package database

import (
	b64 "encoding/base64"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed"
	"github.com/well-informed/wellinformed/graph/model"
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

func buildPage(first int, after *string, edges []*model.Edge) (*model.Connection, error) {
	if after != nil {
		for i := 0; i < len(edges); i++ {
			if *after == edges[i].Cursor {
				if i+1 == len(edges) {
					return nil, errors.New("cursor not found in list")
				} else if i+first+1 > len(edges) {
					edges = edges[i+1:]
				} else {
					edges = edges[i+1 : i+first+1]
				}
				break
			}
		}
	} else if first < len(edges) {
		edges = edges[:first]
	}
	info := &model.PageInfo{
		HasPreviousPage: len(edges) > 0 && after != nil,
		HasNextPage:     len(edges) > first,
		StartCursor:     edges[0].Cursor,
		EndCursor:       edges[len(edges)-1].Cursor,
	}
	return &model.Connection{
		Edges:    edges,
		PageInfo: info,
	}, nil
}

func nodesToEdges(nodes []*model.Node) []*model.Edge {
	edges := make([]*model.Edge, 0)
	for _, node := range nodes {
		edges = append(edges, &model.Edge{
			Node:   node,
			Cursor: b64.StdEncoding.EncodeToString([]byte(string(node.ID))),
		})
	}
	return edges
}
