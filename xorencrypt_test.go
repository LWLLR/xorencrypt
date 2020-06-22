package xorencrypt

import (
	"bufio"
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
	b := make([]rune, 5)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	key := string(b)
	for _, table := range tables {
		rd := strings.NewReader(table.encryptSrc)
		bufRd := bufio.NewReader(rd)

		ioWr := bytes.NewBufferString("")
		bufWr := bufio.NewWriter(ioWr)
		err := Encrypt(key, bufRd, bufWr)
		if err != nil {
			t.Error(err.Error())
		}
		encryptStr := ioWr.String()
		ioWr = bytes.NewBufferString("")
		bufWr = bufio.NewWriter(ioWr)

		err = Encrypt(key, bufio.NewReader(strings.NewReader(encryptStr)), bufWr)
		if err != nil {
			t.Error(err.Error())
		}
		if ioWr.String() != table.encryptSrc {
			t.Errorf("encrypt=>%s decrypt=>%s", table.encryptSrc, ioWr.String())
		}
	}
}
