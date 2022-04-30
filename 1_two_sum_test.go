package leetcode

import (
	"fmt"
	"testing"
)

func Test_twoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	result := twoSum(nums, 9)
	fmt.Printf("%v\n", result)
}
