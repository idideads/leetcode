package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeDuplicates(t *testing.T) {
	nums := []int{1, 2, 2, 3, 3, 3, 5, 7, 11, 13, 13, 17, 19}
	k := removeDuplicates(nums)
	require.Equal(t, 9, k)
}
