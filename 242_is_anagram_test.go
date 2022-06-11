package leetcode

import (
	"fmt"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	s1, s2 := "qwerty", "ytrewq"
	fmt.Println(isAnagram(s1, s2))
}
