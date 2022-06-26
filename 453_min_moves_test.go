package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minMoves(t *testing.T) {
	// fmt.Println(minMoves([]int{1, 10000000}))
	require.Equal(t, 0, minMoves([]int{1, 1, 1}))
	require.Equal(t, 3, minMoves([]int{1, 2, 3}))
	require.Equal(t, 7, minMoves([]int{3, 1, 2, 5}))
}
