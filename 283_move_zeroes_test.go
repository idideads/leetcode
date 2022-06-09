package leetcode

import (
	"fmt"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	nums := []int{0,0,0,5,0,0,3}
	moveZeroes(nums)
	fmt.Println(nums)
}
