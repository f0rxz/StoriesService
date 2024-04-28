package db

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type Database struct {
	db           *sql.DB
	UsersCache   *QueryCache
	StoriesCache *QueryCache
}

func NewDatabase(connectionInfo string) *Database {
	db, err := sql.Open("postgres", connectionInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &Database{
		db:           db,
		UsersCache:   NewQueryCache(1*time.Second, 15*time.Second),
		StoriesCache: NewQueryCache(1*time.Second, 15*time.Second),
	}
}

func (d *Database) Close() error {
	return d.db.Close()
}
