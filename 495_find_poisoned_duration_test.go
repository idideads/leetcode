package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findPoisonedDuration(t *testing.T) {
	require.Equal(t, 0, findPoisonedDuration([]int{1, 3, 5, 7, 9}, 0))
	require.Equal(t, 2, findPoisonedDuration([]int{4}, 2))
	require.Equal(t, 4, findPoisonedDuration([]int{1, 4}, 2))
	require.Equal(t, 17, findPoisonedDuration([]int{1, 7, 10, 14}, 5))
	require.Equal(t, 17, findPoisonedDuration([]int{1, 1, 7, 7, 10, 14}, 5))
}
