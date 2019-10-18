package common

import (
	"crypto/md5"
	"encoding/hex"
	"hash"
)

func GetPasswordMd5(password string) string {
	var (
		h hash.Hash
	)
	h = md5.New()
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}
