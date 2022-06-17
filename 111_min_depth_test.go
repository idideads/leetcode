package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minDepth(t *testing.T) {
	tn5 := &TreeNode{Val: 5}
	tn4 := &TreeNode{Val: 4, Left: tn5}
	tn3 := &TreeNode{Val: 3, Left: tn4}
	tn2 := &TreeNode{Val: 2, Left: tn3}
	tn1 := &TreeNode{Val: 1, Left: tn2}
	
	require.Equal(t, 5, minDepth(tn1))
}
