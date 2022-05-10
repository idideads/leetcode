package leetcode

import (
	"fmt"
	"testing"
)

func Test_isSameTree(t *testing.T) {
	l1 := &TreeNode{Val: 2}
	r1 := &TreeNode{Val: 3}
	root1 := &TreeNode{Val: 1, Left: l1, Right: r1}

	l2 := &TreeNode{Val: 2}
	r2 := &TreeNode{Val: 3}
	root2 := &TreeNode{Val: 1, Left: l2, Right: r2}

	fmt.Println(isSameTree(root1, root2))
}
