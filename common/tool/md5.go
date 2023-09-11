package tool

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5ToString(str string) string {
	data := []byte(str)
	m := md5.New()
	m.Write(data)
	return hex.EncodeToString(m.Sum(nil))
}

func Md5(src []byte) string {
	m := md5.New()
	m.Write(src)
	return hex.EncodeToString(m.Sum(nil))
}
