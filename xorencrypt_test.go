package xorencrypt

import (
	"bytes"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestEncrypt(t *testing.T) {
	tables := []struct {
		encryptSrc string
	}{
		{"test"},
		{"asd.&%#@)(&~@+_)/#。、"},
		{"hello world !"},
	}
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, rand.Intn(len(letterRunes)+1))
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	XOREncryptObj := New(string(b))
	for _, table := range tables {
		rd := strings.NewReader(table.encryptSrc)

		ioWr := bytes.NewBufferString("")
		err := XOREncryptObj.Encrypt(rd, ioWr, nil)
		if err != nil {
			t.Error(err.Error())
		}
		encryptStr := strings.NewReader(ioWr.String())
		ioWr = bytes.NewBufferString("")

		err = XOREncryptObj.Encrypt(encryptStr, ioWr, nil)
		if err != nil {
			t.Error(err.Error())
		}
		if ioWr.String() != table.encryptSrc {
			t.Errorf("encrypt=>%s decrypt=>%s", table.encryptSrc, ioWr.String())
		}
	}
}
