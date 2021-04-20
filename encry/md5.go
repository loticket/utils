package encry

import (
	"crypto/md5"
	"encoding/hex"
)

//字符串md5加密
func Md5Str(plainText string) string {
	return MD5([]byte(plainText))
}

// MD5 MD5哈希加密， 返回32位字符串
func MD5(plainText []byte) string {
	m := md5.New()
	m.Write(plainText)
	return hex.EncodeToString(m.Sum(nil))
}
