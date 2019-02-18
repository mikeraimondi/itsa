package itsa

import (
	"fmt"

	"github.com/mikeraimondi/skiplist"
)

const maxSkipListLevel = 1024

func newSkiplistMT() (*skiplistMT, error) {
	list, err := skiplist.New(maxSkipListLevel)
	if err != nil {
		return nil, fmt.Errorf("initializing skiplist: %s", err)
	}
	return &skiplistMT{list}, nil
}

type skiplistMT struct {
	list *skiplist.List
}

func (s *skiplistMT) put(key, val []byte) {
	s.list.Insert(key, val)
}

func (s *skiplistMT) get(key []byte) ([]byte, bool) {
	return s.list.Search(key)
}

func (s *skiplistMT) flush() (*sstable, error) {
	// write memtable to disk, creating new sstable
	return nil, nil
}
