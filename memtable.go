package lsmdb

type memtable struct {
	// skiplist
}

func (m *memtable) flush() (*sstable, error) {
	// write memtable to disk, creating new sstable
	return nil, nil
}
