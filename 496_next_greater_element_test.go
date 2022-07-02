package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_nextGreaterElement(t *testing.T) {
	require.EqualValues(t, []int{-1, 3, -1}, nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	require.EqualValues(t, []int{3, -1}, nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}
