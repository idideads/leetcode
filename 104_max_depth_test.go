package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxDepth(t *testing.T) {
	tn5 := &TreeNode{Val: 5}
	tn4 := &TreeNode{Val: 4, Right: tn5}
	tn2 := &TreeNode{Val: 2, Left: tn4}
	tn3 := &TreeNode{Val: 3}
	tn1 := &TreeNode{Val: 1, Left: tn2, Right: tn3}
	require.Equal(t, 4, maxDepth(tn1))
}
