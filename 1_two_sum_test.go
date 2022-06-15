package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_twoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	require.EqualValues(t, []int{0, 1}, twoSum(nums, 9))
}
