package util

import (
	"crypto/md5"
	"encoding/hex"
)

/**
md5加密
*/
func Md5encode(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}
