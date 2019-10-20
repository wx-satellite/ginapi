package common

import (
	"crypto/md5"
	"encoding/hex"
	"gopkg.in/go-playground/validator.v9"
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

func Validate(model interface{}) error {
	var (
		v *validator.Validate
	)
	v = validator.New()
	return v.Struct(model)
}
