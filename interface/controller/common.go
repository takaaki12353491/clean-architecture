package controller

import (
	"math/rand"
	"time"
)

const (
	letters   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	indexBit  = 6
	indexMask = 1<<indexBit - 1
	indexMax  = 63 / indexBit
)

func createRand() (randVal string) {
	randSource := rand.NewSource(time.Now().UnixNano())
	n := 32
	b := make([]byte, n)
	cache, remain := randSource.Int63(), indexMax
	for i := n - 1; i >= 0; {
		if remain == 0 {
			cache, remain = randSource.Int63(), indexMax
		}
		index := int(cache & indexMask)
		if index < len(letters) {
			b[i] = letters[index]
			i--
		}
		cache >>= indexBit
		remain--
	}
	randVal = string(b)
	return
}
