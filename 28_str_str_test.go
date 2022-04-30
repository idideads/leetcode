package leetcode

import (
	"fmt"
	"testing"
)

func Test_strStr(t *testing.T) {
	haystack, needle := "aa", "aa"
	startPosition := strStr(haystack, needle)
	fmt.Printf("startPosition is %d\n", startPosition)
}
