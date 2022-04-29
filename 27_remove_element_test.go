package leetcode

import (
	"fmt"
	"testing"
)

func Test_removeElement(t *testing.T) {
	nums := []int{1, 2, 3, 3, 5, 7, 3, 11}
	for _, n := range nums {
		fmt.Printf("%d\t", n)
	}
	p := removeElement(nums, 3)
	fmt.Printf("\nK is %d\n", p)
	for _, n := range nums {
		fmt.Printf("%d\t", n)
	}
}
