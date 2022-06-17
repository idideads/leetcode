package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sortedArrayToBST(t *testing.T) {
	nums := []int{1, 2, 3, 5, 7, 11, 13, 17}
	root := sortedArrayToBST(nums)

	require.NotNil(t, root)
	require.EqualValues(t, nums, inorderTraversal(root))
}
