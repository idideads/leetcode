package leetcode

import (
	"fmt"
	"testing"
)

func Test_postorderTraversal(t *testing.T) {
	tn5 := &TreeNode{Val: 5}
	tn4 := &TreeNode{Val: 4}
	tn2 := &TreeNode{Val: 2, Left: tn4}
	tn3 := &TreeNode{Val: 3, Right: tn5}
	tn1 := &TreeNode{Val: 1, Left: tn2, Right: tn3}
	fmt.Println(postorderTraversal(tn1))
}
