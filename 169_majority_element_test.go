package leetcode

import (
	"fmt"
	"testing"
)

func Test_majorityElement(t *testing.T) {
	nums := []int{2, 2, 2, 11, 2, 2, 1, 3, 4, 2, 2, 2}
	fmt.Println(majorityElement(nums))
}
