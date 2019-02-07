package lsmdb

func NewDB() (*DB, error) {
	return nil, nil
}

type DB struct {
	// memtable
}

func (db *DB) NewIterator() *Iterator {
	return nil
}

func (db *DB) Put(key, value []byte) error {
	return nil
}

// func (db *DB) Write(keys, values) error

func (db *DB) Get(key []byte) ([]byte, error) {
	return nil, nil
}

// func (db *DB) MultiGet(keys []byte...) ([]bool, error)

func (db *DB) Delete(key []byte) error {
	return nil
}
