package xorencrypt

import (
	"io"
)

//XOREncrypt init Data
type XOREncrypt struct {
	key string
}

//SetKey 设置加密字符串
func (x *XOREncrypt) SetKey(key string) {
	x.key = key
}

//GetKey 获取加密字符串
func (x *XOREncrypt) GetKey() string {
	return x.key
}

//New new obj
func New(key string) *XOREncrypt {
	return &XOREncrypt{
		key: key,
	}
}

//Encrypt 加密
func (x *XOREncrypt) Encrypt(rd io.Reader, wr io.Writer, buf []byte) error {
	secureKey := []byte(x.key)

	maxLen := len(secureKey)
	curIndex := 0
	var block []byte
	if buf == nil {
		block = make([]byte, 4096)
	} else {
		block = buf
	}
	var tmp byte
	for {
		n, err := rd.Read(block)
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
		wr.Write(block[:n])
	}
	return nil
}
