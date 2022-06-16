package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_plusOne(t *testing.T) {
	require.EqualValues(t, []int{1, 2, 4}, plusOne([]int{1, 2, 3}))
	require.EqualValues(t, []int{3, 0, 0}, plusOne([]int{2, 9, 9}))
	require.EqualValues(t, []int{1}, plusOne([]int{0}))
}
