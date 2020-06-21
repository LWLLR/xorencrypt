package xorencrypt

import (
	"bufio"
	"io"
)

//Encrypt 加密
func Encrypt(key string, buf *bufio.Reader, encryptFileBuf *bufio.Writer) error {
	secureKey := []byte(key)

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
	encryptFileBuf.Flush()
	return nil
}
