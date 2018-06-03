package tcloud

import (
	"math/rand"
	"crypto/md5"
	"encoding/hex"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func GetRandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func MD5(in string) (out string) {
	h := md5.New()
	h.Write([]byte(in))
	out = hex.EncodeToString(h.Sum(nil))
	return
}
