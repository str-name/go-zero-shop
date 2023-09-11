package tool

import "math/rand"

const (
	randAll    = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	randLetter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	randNum    = "0123456789"
)

const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1

	numIdxBits = 4
	numIdxMask = 1<<numIdxBits - 1
)

func RandAllToString(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(randAll) {
			b[i] = randAll[idx]
			i++
		}
	}
	return string(b)
}

func RandLetterToString(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(randLetter) {
			b[i] = randLetter[idx]
			i++
		}
	}
	return string(b)
}

func RandNumToString(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & numIdxMask); idx < len(randNum) {
			b[i] = randNum[idx]
			i++
		}
	}
	return string(b)
}
