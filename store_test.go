package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestPathTransformFunc (t *testing.T) {
	key := "bestpictures"
	pathKey := CASPathTransformFunc(key)
	expectedOriginal := "b2f8f1dd50fdeec113ac1b5066d1d3a10f70f1dc";
	expectedPathName := "b2f8f/1dd50/fdeec/113ac/1b506/6d1d3/a10f7/0f1dc";
	if pathKey.PathName != expectedPathName {
		t.Errorf("have %s expected %s\n", pathKey.PathName, expectedPathName)
	}
	if pathKey.Filename != expectedOriginal {
		t.Errorf("have %s expected %s\n", pathKey.Filename, expectedOriginal)
	}
}

func TestDelete (t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	id := generateID()
	key := "mypicture"
	data := []byte("some jpg bytes")
	if _, err := s.writeStream(id, key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}
	if err := s.Delete(id, key); err != nil {
		t.Error(err)
	}
}

func TestStore(t *testing.T) {
	s := newStore()
	id := generateID()
	defer tearDown(t, s)

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("foo_%d", i)
		data := []byte("some jpg bytes")
		if _, err := s.writeStream(id, key, bytes.NewReader(data)); err != nil {
			t.Error(err)
		}

		if ok := s.Has(id, key); !ok {
			t.Error("expected to have key", key)
		}
		
		_, r, err := s.Read(id, key)
		if err != nil {
			t.Error(err)
		}

		b, _ := io.ReadAll(r)
		if string(b) != string(data) {
			t.Errorf("expected %s got %s", data, b)
		}

		if err := s.Delete(id, key); err != nil {
			t.Error(err)
		}

		if ok := s.Has(id, key); ok {
			t.Errorf("expected to not have key %s", key)
		}
	}
}

func newStore () *Store {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	return NewStore(opts)
}

func tearDown (t *testing.T, s *Store) {
	if err := s.Clear(); err != nil {
		t.Error(err)
	}
}