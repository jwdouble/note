package auth

import (
	"fmt"
	"testing"
)

// ase  cbc pkcs5  128

func TestEncryptCBC(t *testing.T) {
	src := []byte("12345678")
	key := []byte("1122334455667788")
	iv := []byte("1122334455667788")
	fmt.Println(EncryptCBC(src, key, iv))
}
