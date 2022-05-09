package leetcode

import (
	"fmt"
	"testing"
)

func Test_merge(t *testing.T) {
	numbs1 := []int{4, 5, 6, 0, 0, 0}
	numbs2 := []int{1, 2, 3}
	m, n := 3, 3
	merge(numbs1, m, numbs2, n)
	fmt.Println(numbs1)
}
