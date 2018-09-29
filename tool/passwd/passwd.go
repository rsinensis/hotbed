package passwd

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
)

const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func getRandomString(n int, alphabets ...byte) string {

	var bytes = make([]byte, n)

	rand.Read(bytes)
	for i, b := range bytes {
		if len(alphabets) == 0 {
			bytes[i] = alphanum[b%byte(len(alphanum))]
		} else {
			bytes[i] = alphabets[b%byte(len(alphabets))]
		}
	}

	return string(bytes)
}

func GetSalt() string {

	return getRandomString(10)

}

func GetPassword(str string) string {

	h := md5.New()
	h.Write([]byte(str))
	cipherStr := h.Sum(nil)

	return hex.EncodeToString(cipherStr)
}
