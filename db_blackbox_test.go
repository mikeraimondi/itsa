package lsmdb_test

import (
	"bytes"
	"testing"

	"github.com/mikeraimondi/lsmdb"
)

func TestPutAndGet(t *testing.T) {
	db, err := lsmdb.New()
	if err != nil {
		t.Fatalf("expected no error from New. got %s", err)
	}

	key := []byte("foo")
	expectedVal := []byte("bar")
	err = db.Put(key, expectedVal)
	if err != nil {
		t.Fatalf("expected no error from Put. got %s", err)
	}

	val, err := db.Get(key)
	if err != nil {
		t.Fatalf("expected no error from Get. got %s", err)
	}
	if bytes.Compare(expectedVal, val) != 0 {
		t.Fatalf("wrong value from Get. expected %q. got %q", expectedVal, val)
	}
}
