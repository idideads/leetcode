package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isBalanced(t *testing.T) {
	tn2 := &TreeNode{Val: 2}
	tn3 := &TreeNode{Val: 3}
	tn1 := &TreeNode{Val: 1, Left: tn2, Right: tn3}
	require.Equal(t, true, isBalanced(tn1))
	require.Equal(t, true, isBalanced(tn2))
	require.Equal(t, true, isBalanced(tn3))
}
