package leetcode

import "fmt"

func hammingDistance(x int, y int) int {
	binarySlice := fmt.Sprintf("%b", x^y)
	n := 0
	for _, b := range binarySlice {
		if b == '1' {
			n++
		}
	}
	return n
}
