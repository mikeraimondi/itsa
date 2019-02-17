package itsa

import (
	"fmt"
)

func New() (*DB, error) {
	memtable, err := newSkiplistMT()
	if err != nil {
		return nil, fmt.Errorf("initializing memtable: %s", err)
	}

	db := &DB{
		mt: memtable,
	}
	return db, nil
}

type DB struct {
	mt memtable
}

func (db *DB) NewIterator() *Iterator {
	return nil
}

func (db *DB) Put(key, value []byte) error {
	db.mt.put(key, value)

	return nil
}

// func (db *DB) Write(keys, values) error

func (db *DB) Get(key []byte) ([]byte, error) {
	result, _ := db.mt.get(key)

	return result, nil
}

// func (db *DB) MultiGet(keys []byte...) ([]bool, error)

func (db *DB) Delete(key []byte) error {
	return nil
}

type memtable interface {
	put([]byte, []byte)
	get([]byte) ([]byte, bool)
	flush() (*sstable, error)
}
