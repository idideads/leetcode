package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findRelativeRanks(t *testing.T) {
	require.EqualValues(t, []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}, findRelativeRanks([]int{5, 4, 3, 2, 1}))
	require.EqualValues(t, []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"}, findRelativeRanks([]int{10, 3, 8, 9, 4}))
}
