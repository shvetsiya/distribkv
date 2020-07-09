package db

import (
	bolt "go.etcd.io/bbolt"
)

// Database is an open bolt DB
type Database struct {
	db *bolt.DB
}

// NewDatabase returns an instance of DB that we can work with
func NewDatabase(dbPath string) (db *Database, closeFunc func() error, err error) {
	boltDb, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, nil, err
	}
	closeFunc = boltDb.Close
	return &Database{db: boltDb}, closeFunc, nil
}
