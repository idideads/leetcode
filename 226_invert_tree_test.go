package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_invertTree(t *testing.T) {
	root := &TreeNode{Val: 4,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}},
	}
	root = invertTree(root)
	require.NotNil(t, root)
	require.EqualValues(t, []int{4, 7, 9, 6, 2, 3, 1}, preorderTraversal(root))
	root = invertTree(root)
	require.NotNil(t, root)
	require.EqualValues(t, []int{4, 2, 1, 3, 7, 6, 9}, preorderTraversal(root))
}
