package xorencrypt

import (
	"io"
)

//XOREncrypt init Data
type XOREncrypt struct {
	key string
}

//NewXOREncrypt new obj
func NewXOREncrypt(key string) *XOREncrypt {
	return &XOREncrypt{
		key: key,
	}
}

//Encrypt 加密
func (x *XOREncrypt) Encrypt(buf io.Reader, encryptFileBuf io.Writer) error {
	secureKey := []byte(x.key)

	maxLen := len(secureKey)
	curIndex := 0
	block := make([]byte, 4096)
	var tmp byte
	for {
		n, err := buf.Read(block)
		if err != nil && err != io.EOF {
			return err
		}
		if 0 == n || err == io.EOF {
			break
		}
		for i := 0; i < n; i++ {
			tmp = block[i]
			block[i] = tmp ^ secureKey[curIndex]
			curIndex++
			curIndex = curIndex % maxLen
		}
		encryptFileBuf.Write(block[:n])
	}
	return nil
}
