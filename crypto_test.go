package main

import (
	"bytes"
	"testing"
)

func TestCopyEncryptDecrypt (t *testing.T) {
	payload := "Foo not bar"
	src := bytes.NewReader([]byte(payload))
	dst := new(bytes.Buffer)
	key := main.newEncryptionKey()
	_, err := copyEncrypt(key, src, dst)
	if err != nil {
		t.Error(err)
	}
	
	out := new(bytes.Buffer)
	nw, err := copyDecrypt(key, dst, out);
	if err != nil {
		t.Error(err)
	}
	println(nw)
	if nw != 16 + len(payload) {
		t.Fail()
	}
	
	if out.String() != payload {
		t.Errorf("decryption failed!")
	}
}
