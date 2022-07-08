package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_binaryTreePaths(t *testing.T) {
	tn4 := &TreeNode{Val: 4}
	tn2 := &TreeNode{Val: 2}
	tn3 := &TreeNode{Val: 3, Right: tn4}
	tn1 := &TreeNode{Val: 1, Left: tn2, Right: tn3}
	require.EqualValues(t, []string{"1->2", "1->3->4"}, binaryTreePaths(tn1))
}
