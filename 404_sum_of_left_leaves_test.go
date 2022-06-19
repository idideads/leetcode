package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sumOfLeftLeaves(t *testing.T) {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	require.Equal(t, 24, sumOfLeftLeaves(root))

	root = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}}
	require.Equal(t, 10, sumOfLeftLeaves(root))

}
