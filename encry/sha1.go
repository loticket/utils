package encry

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1String(plainText string) string {
	return Sha1([]byte(plainText))
}

// SHA1 SHA1哈希加密
func Sha1(plainText []byte) string {
	sha := sha1.New()
	sha.Write(plainText)
	return hex.EncodeToString(sha.Sum(nil))
}
