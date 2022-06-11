package leetcode

import (
	"fmt"
	"testing"
)

func Test_wordPattern(t *testing.T) {
	pattern, str := "abba", "dog dog dog dog"
	fmt.Println(wordPattern(pattern, str))
}
