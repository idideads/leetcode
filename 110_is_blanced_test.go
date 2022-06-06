package leetcode

import (
	"fmt"
	"testing"
)

func Test_isBalanced(t *testing.T) {
	tn2 := &TreeNode{Val: 2}
	tn3 := &TreeNode{Val: 3}
	tn1 := &TreeNode{Val: 1, Left: tn2, Right: tn3}
	fmt.Println(isBalanced(tn1))
}
