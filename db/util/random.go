package util

import (
	"math/rand"
	"strings"
	"time"
)

const alpabets = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int32) int32 {
	return rand.Int31n(max - min + 1)
}

func RandomString(n int) string {
	var name strings.Builder
	l := len(alpabets)
	for i := 0; i < n; i++ {
		char := alpabets[rand.Intn(l)]
		name.WriteByte(char)
	}

	return name.String()
}

func RandomName(n int) string {
	return RandomString(n)
}

func RandomAmount(min, max int32) int32 {
	return RandomInt(min, max)
}

func RandomAge(min, max int32) int32 {
	return RandomInt(min, max)
}

func RandomCurrency() string {
	currs := []string{"EUR", "USD", "NGN"}
	return currs[rand.Intn(len(currs))]
}
