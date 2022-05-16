package auth

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// usually use base64 encoder make password readable
// key must 16 length
// iv must equal to block size 16

func EncryptCBC(src, key, iv []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	ret := func(cipherText []byte, blockSize int) []byte { //PKCS5Padding
		padding := blockSize - len(cipherText)%blockSize
		padText := bytes.Repeat([]byte{byte(padding)}, padding)
		return append(cipherText, padText...)
	}(src, block.BlockSize())

	blockMode := cipher.NewCBCEncrypter(block, iv)
	buf := make([]byte, len(ret)) // 以2的指数值舍弃多余的容量（8，16，32)
	copy(buf, ret)
	blockMode.CryptBlocks(buf, buf)

	return base64.StdEncoding.EncodeToString(buf)
}
