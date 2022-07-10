package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_perorderTravelMap(t *testing.T) {
	root := &TreeNode{Val: 20,
		Left: &TreeNode{Val: 10,
			Left:  &TreeNode{Val: 10},
			Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 10, Right: &TreeNode{Val: 20}}}},
		Right: &TreeNode{Val: 30},
	}
	require.EqualValues(t, map[int]int{10: 3, 20: 3, 30: 1}, perorderTravelMap(root))
}

func Test_findMode(t *testing.T) {
	root := &TreeNode{Val: 20,
		Left: &TreeNode{Val: 10,
			Left:  &TreeNode{Val: 10},
			Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 10, Right: &TreeNode{Val: 20}}}},
		Right: &TreeNode{Val: 30},
	}
	require.EqualValues(t, []int{20, 10}, findMode(root))
}
