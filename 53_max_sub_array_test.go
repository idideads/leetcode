package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	require.Equal(t, 6, maxSubArray(nums))
}
