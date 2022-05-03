package leetcode

import (
	"fmt"
	"testing"
)

func Test_searchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 7
	fmt.Printf("insert position:%d\n", searchInsert(nums, target))
}
