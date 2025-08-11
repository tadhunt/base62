package base62

import (
	"bytes"
	"crypto/rand"
	"fmt"
)

func NewBase62Id(length int) (string, error) {
	const base62chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	b := make([]byte, length)

	n, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	if n != len(b) {
		return "", fmt.Errorf("short read got %d expected %d", n, len(b))
	}

	var buf bytes.Buffer
	for _, v := range b {
		buf.WriteByte(base62chars[int(v)%len(base62chars)])
	}

	return buf.String(), nil
}
