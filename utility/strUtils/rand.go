package strUtils

import (
	"math/rand"
	"time"
)

const (
	letter = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func Rand(n int) string {
	byteStr := []byte(letter)
	randStr := make([]byte, 0, n)
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < n; i++ {
		randStr = append(randStr, byteStr[rand.Intn(len(byteStr))])
	}
	return string(randStr)
}
