package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_constructRectangle(t *testing.T) {
	require.EqualValues(t, []int{2, 2}, constructRectangle(4))
	require.EqualValues(t, []int{37, 1}, constructRectangle(37))
	require.EqualValues(t, []int{10, 6}, constructRectangle(60))
	require.EqualValues(t, []int{427, 286}, constructRectangle(122122))
}
