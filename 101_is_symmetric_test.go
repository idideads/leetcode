package leetcode

import (
	"fmt"
	"testing"
)

func Test_isSymmetric(t *testing.T) {
	ltn5 := &TreeNode{Val: 5}
	ltn6 := &TreeNode{Val: 6}
	ltn7 := &TreeNode{Val: 7}
	ltn8 := &TreeNode{Val: 8}

	ltn3 := &TreeNode{Val: 3, Left: ltn5, Right: ltn6}
	ltn4 := &TreeNode{Val: 4, Left: ltn7, Right: ltn8}
	ltn2 := &TreeNode{Val: 2, Left: ltn3, Right: ltn4}

	rtn5 := &TreeNode{Val: 5}
	rtn6 := &TreeNode{Val: 6}
	rtn7 := &TreeNode{Val: 7}
	rtn8 := &TreeNode{Val: 8}

	rtn3 := &TreeNode{Val: 3, Left: rtn6, Right: rtn5}
	rtn4 := &TreeNode{Val: 4, Left: rtn8, Right: rtn7}
	rtn2 := &TreeNode{Val: 2, Left: rtn4, Right: rtn3}

	tn1 := &TreeNode{Val: 1, Left: ltn2, Right: rtn2}

	fmt.Println(isSymmetric(tn1))
}

func Test_isMirror(t *testing.T) {
	ltn5 := &TreeNode{Val: 5}
	ltn6 := &TreeNode{Val: 6}
	ltn7 := &TreeNode{Val: 7}
	ltn8 := &TreeNode{Val: 8}

	ltn3 := &TreeNode{Val: 3, Left: ltn5, Right: ltn6}
	ltn4 := &TreeNode{Val: 4, Left: ltn7, Right: ltn8}
	ltn2 := &TreeNode{Val: 2, Left: ltn3, Right: ltn4}

	rtn5 := &TreeNode{Val: 5}
	rtn6 := &TreeNode{Val: 6}
	rtn7 := &TreeNode{Val: 7}
	rtn8 := &TreeNode{Val: 8}

	rtn3 := &TreeNode{Val: 3, Left: rtn6, Right: rtn5}
	rtn4 := &TreeNode{Val: 4, Left: rtn8, Right: rtn7}
	rtn2 := &TreeNode{Val: 2, Left: rtn4, Right: rtn3}

	fmt.Println(isMirror(ltn2, rtn2))
}
