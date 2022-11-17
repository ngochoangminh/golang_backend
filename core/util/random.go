package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomStr(n int) string {
	var sb strings.Builder
	alphabet := "abcdefghijklmnopqrstvxywz"
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomStr(6)
}

func RandomMoney() int64 {
	return RandomInt(1, 100000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "GBP", "CAD", "UER", "VND"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func NanoId(n int) string {
	var sb strings.Builder
	alphabet := "0123456789abcdefghijklmnopqrstvxywzABCDEFGHIJKLMNOPQRSTVXYWZ"
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}
