package leetcode

import (
	"fmt"
	"testing"
)

func Test_containsNearbyDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	k := 3
	fmt.Println(containsNearbyDuplicate(nums, k))
}
