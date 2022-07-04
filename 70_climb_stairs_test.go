package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_climbStairs(t *testing.T) {
	require.Equal(t, 2, climbStairs(2))
	require.Equal(t, 3, climbStairs(3))
	require.Equal(t, 5, climbStairs(4))
	require.Equal(t, 1836311903, climbStairs(45))
}
