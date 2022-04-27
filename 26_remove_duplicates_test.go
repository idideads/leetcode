package leetcode

import (
	"fmt"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	nums := []int{1,2,2,3,3,3,5,7,11,13,13,17,19}
	k := removeDuplicates(nums)
	fmt.Printf("K is %d\n", k)
	for _, n := range nums {
		fmt.Printf("%d\t", n)
	}
}
