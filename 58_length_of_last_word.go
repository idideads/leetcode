package leetcode

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	list := strings.Split(strings.Trim(s, " "), " ")
	size := len(list)
	return len(list[size-1])
}
