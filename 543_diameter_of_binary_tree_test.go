package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}
	require.Equal(t, 3, diameterOfBinaryTree(root))
}
