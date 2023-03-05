package utilities

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

func GetMD5Hash(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

func GetSHA1Hash(s string) string {
	hash := sha1.Sum([]byte(s))

	return hex.EncodeToString(hash[:])
}

func GetSHA256Hash(s string) string {
	hash := sha256.Sum256([]byte(s))

	return hex.EncodeToString(hash[:])
}

func GetSHA512Hash(s string) string {
	hash := sha512.Sum512([]byte(s))

	return hex.EncodeToString(hash[:])
}

func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Base64Decode(s string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}
