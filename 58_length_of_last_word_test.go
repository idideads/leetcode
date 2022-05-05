package leetcode

import (
	"fmt"
	"testing"
)

func Test_lengthOfLastWord(t *testing.T) {
	str := "hello  world  "
	fmt.Printf("lastword len = %d\n", lengthOfLastWord(str))
}
