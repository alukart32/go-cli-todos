package boltdbx

import (
	"log"
	"sync"
	"time"

	"github.com/boltdb/bolt"
)

var (
	db       *bolt.DB
	_dbName  = "./pkg/boltdbx/todos.db"
	_bucket  = "todos"
	_options = &bolt.Options{Timeout: 1 * time.Second}
	_once    sync.Once
)

func New() *bolt.DB {
	_once.Do(func() {
		var err error
		db, err = bolt.Open(_dbName, 0600, _options)
		if err != nil {
			log.Fatalf("open the boltdb: %s", err)
		}
		err = db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(_bucket))
			return err
		})
		if err != nil {
			log.Fatalf("bucket %s: %s", _bucket, err)
		}
	})
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}
