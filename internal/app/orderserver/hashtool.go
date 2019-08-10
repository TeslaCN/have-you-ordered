package orderserver

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5FromByte(b []byte) string {
	hash := md5.New()
	hash.Write(b)
	return hex.EncodeToString(hash.Sum(nil))
}

func Md5FromString(s string) string {
	return Md5FromByte([]byte(s))
}
