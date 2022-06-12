package leetcode

import (
	"fmt"
	"testing"
)

func Test_reverseString(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	for _, v := range s {
		fmt.Printf("%c\t", v)
	}
}
