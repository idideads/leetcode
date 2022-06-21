package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_thirdMax(t *testing.T) {
	require.Equal(t, 1, thirdMax([]int{2, 2, 3, 1}))
	require.Equal(t, 3, thirdMax([]int{3, 2, 2, 2}))
	require.Equal(t, 7, thirdMax([]int{11, 5, 13, 7, 7, 7, 5, 5, 11, 11, 11, 13}))
}
