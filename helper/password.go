package helper

import (
	"crypto/md5"
	"encoding/hex"
)

func HashPassword(plainText string) string {
	hash := md5.Sum([]byte(plainText))
	return hex.EncodeToString(hash[:])
}
