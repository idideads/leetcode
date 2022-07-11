package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findLHS(t *testing.T) {
	require.Equal(t, 5, findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
	require.Equal(t, 0, findLHS([]int{1, 3, 5, 7, 9}))
	require.Equal(t, 0, findLHS([]int{3, 3, 3, 3, 3}))
}
