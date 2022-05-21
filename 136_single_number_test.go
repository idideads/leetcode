package leetcode

import (
	"fmt"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	nums := []int{4, 1, 2, 1, 2, 6, 7, 6, 7}
	fmt.Println(singleNumber(nums))
}
