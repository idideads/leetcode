package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxProfit(t *testing.T) {
	input := []int{7, 1, 5, 3, 6, 4}
	require.Equal(t, 5, maxProfit(input))

	input = []int{7, 6, 5, 4, 3}
	require.Equal(t, 0, maxProfit(input))

	input = []int{2, 4, 1}
	require.Equal(t, 2, maxProfit(input))
}
