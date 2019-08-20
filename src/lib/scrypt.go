package lib

import (
	"encoding/base64"

	"golang.org/x/crypto/scrypt"
)

// Scrypt 用于加密
func Scrypt(pass string) string {
	salt := []byte{0xc8, 0x28, 0xf2, 0x58, 0xa7, 0x6a, 0xad, 0x7b}

	dk, error := scrypt.Key([]byte(pass), salt, 1<<15, 8, 1, 32)
	if error != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(dk)
}
