package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_moveZeroes(t *testing.T) {
	nums := []int{0, 0, 0, 5, 0, 0, 3}
	moveZeroes(nums)
	require.EqualValues(t, []int{5, 3, 0, 0, 0, 0, 0}, nums)
}
