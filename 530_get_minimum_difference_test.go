package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getMinimumDifference(t *testing.T) {
	root := &TreeNode{Val: 9, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 48, Left: &TreeNode{Val: 12}, Right: &TreeNode{Val: 58}}}
	require.Equal(t, 3, getMinimumDifference(root))
}
