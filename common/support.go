package common

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
	"time"
)

func RandomCharacter() string {
	return RandomString(1)
}

func RandomString(len int) string {
	buff := make([]byte, len)
	_, _ = rand.Read(buff)
	str := base64.StdEncoding.EncodeToString(buff)
	return str[:len]
}

func KeepAlive() {
	for {
		time.Sleep(time.Second)
	}
}

func IsNullOrWhitespace(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
