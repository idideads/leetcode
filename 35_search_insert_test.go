package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_searchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	require.Equal(t, 2, searchInsert(nums, 5))
	require.Equal(t, 1, searchInsert(nums, 2))
	require.Equal(t, 4, searchInsert(nums, 7))
}
