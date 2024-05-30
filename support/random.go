package support

import (
	"crypto/rand"
	"encoding/base64"
)

func RandomString(len int) string {
	buff := make([]byte, len)
	_, _ = rand.Read(buff)
	str := base64.StdEncoding.EncodeToString(buff)
	return str[:len]
}
