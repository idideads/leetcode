package leetcode

import (
	"fmt"
	// "log"
	// "reflect"
	"testing"
)

func Test_shortestToChar(t *testing.T) {
	answer1 := shortestToChar("baab", 'b')
	answer2 := shortestToChar("loveleetcodeabc", 'e')
	fmt.Printf("%v\n", answer1)
	fmt.Printf("%v\n", answer2)
}
