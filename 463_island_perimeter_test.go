package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_islandPerimeter(t *testing.T) {
	grid := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
	require.Equal(t, 16, islandPerimeter(grid))

	require.Equal(t, 4, islandPerimeter([][]int{{1}}))
	require.Equal(t, 4, islandPerimeter([][]int{{1, 0}}))
}
