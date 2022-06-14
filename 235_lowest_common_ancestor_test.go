package leetcode

import (
	"fmt"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	tn3 := &TreeNode{Val: 3}
	tn5 := &TreeNode{Val: 5}
	tn4 := &TreeNode{Val: 4, Left: tn3, Right: tn5}
	tn0 := &TreeNode{Val: 0}
	tn7 := &TreeNode{Val: 7}
	tn9 := &TreeNode{Val: 9}
	tn2 := &TreeNode{Val: 2, Left: tn0, Right: tn4}
	tn8 := &TreeNode{Val: 8, Left: tn7, Right: tn9}
	tn6 := &TreeNode{Val: 6, Left: tn2, Right: tn8}
	fmt.Println(lowestCommonAncestor(tn6, tn0, tn9))
	fmt.Println(lowestCommonAncestor(tn6, tn0, tn5))
}
