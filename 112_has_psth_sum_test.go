package leetcode

import (
	"fmt"
	"testing"
)

func Test_hasPathSum(t *testing.T) {
	tn4 := &TreeNode{Val: 4}
	tn2 := &TreeNode{Val: 2}
	tn3 := &TreeNode{Val: 3, Right: tn4}
	tn1 := &TreeNode{Val: 1, Left: tn2, Right: tn3}
	fmt.Println(hasPathSum(tn1, 8))
}
