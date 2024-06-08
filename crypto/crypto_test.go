package crypto

import (
	"bytes"
	"testing"
)

func TestCopyEncryptDecrypt (t *testing.T) {
	payload := "Foo not bar"
	src := bytes.NewReader([]byte(payload))
	dst := new(bytes.Buffer)
	key := NewEncryptionKey()
	_, err := CopyEncrypt(key, src, dst)
	if err != nil {
		t.Error(err)
	}
	
	out := new(bytes.Buffer)
	nw, err := CopyDecrypt(key, dst, out);
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
