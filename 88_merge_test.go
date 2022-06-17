package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_merge(t *testing.T) {
	numbs1 := []int{4, 5, 6, 0, 0, 0}
	numbs2 := []int{1, 2, 3}
	m, n := 3, 3
	merge(numbs1, m, numbs2, n)
	require.NotNil(t, numbs1)
	require.EqualValues(t, []int{1, 2, 3, 4, 5, 6}, numbs1)
}
