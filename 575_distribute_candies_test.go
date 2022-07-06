package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_distributeCandies(t *testing.T) {
	require.Equal(t, 3, distributeCandies([]int{1, 1, 2, 2, 3, 3}))
	require.Equal(t, 1, distributeCandies([]int{7, 7, 7, 7, 7}))
}
