package utilities

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

func GetSha1Hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)

	return hex.EncodeToString(bs)
}

func GetMd5Hash(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}
