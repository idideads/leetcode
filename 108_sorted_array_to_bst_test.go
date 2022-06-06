package leetcode

import (
	"fmt"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	nums := []int{1, 2, 3, 5, 7, 11, 13, 17}
	root := sortedArrayToBST(nums)
	fmt.Println(inorderTraversal(root))
}
