package tool

import (
	"fmt"
	"testing"
)

const (
	RandAll    = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	RandLetter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	RandNum    = "0123456789"
)

func TestRandAllToString(t *testing.T) {
	fmt.Println(RandAllToString(12))
	fmt.Println(RandLetterToString(12))
	fmt.Println(RandNumToString(12))
}
