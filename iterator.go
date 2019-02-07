package lsmdb

type Iterator struct{}

func (i *Iterator) Seek(key []byte) error {
	return nil
}

func (i *Iterator) Next() {
}

func (i *Iterator) Key() {
}

func (i *Iterator) Value() {
}
