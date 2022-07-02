package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMaxConsecutiveOnes(t *testing.T) {
	require.Equal(t, 0, findMaxConsecutiveOnes([]int{0}))
	require.Equal(t, 3, findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
	require.Equal(t, 2, findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}))
}
