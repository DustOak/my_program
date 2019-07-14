package tools

import (
	"crypto/rand"
)

func GetSid() string {
	value, err := rand.Prime(rand.Reader, 64)
	if err != nil {
		panic(err)
	}
	return value.String()
}
