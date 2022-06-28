package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findContentChildren(t *testing.T) {
	require.Equal(t, 0, findContentChildren([]int{1, 2, 3}, []int{}))
	require.Equal(t, 1, findContentChildren([]int{1, 2, 3}, []int{1, 1}))
	require.Equal(t, 2, findContentChildren([]int{1, 2}, []int{1, 2, 3}))
	require.Equal(t, 3, findContentChildren([]int{2, 4, 6, 8}, []int{1, 3, 5, 7}))
}
