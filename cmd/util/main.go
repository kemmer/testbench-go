package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomString(l int) string {
	charset := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" +
		" _-"

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano())) //nolint:gosec

	b := make([]byte, l)

	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}

func main() {
	fmt.Println("randomString:", randomString(99))
}
