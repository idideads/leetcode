package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_summaryRanges(t *testing.T) {
	nums := []int{0, 2, 3, 4, 6, 8, 9}
	require.NotNil(t, summaryRanges(nums))
	require.EqualValues(t, []string{"0", "2->4", "6", "8->9"}, summaryRanges(nums))
}
