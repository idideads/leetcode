package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_singleNumber(t *testing.T) {
	nums := []int{4, 1, 2, 1, 2, 6, 7, 6, 7}
	require.Equal(t, 4, singleNumber(nums))
}

func Test_singleNumber2(t *testing.T) {
	nums := []int{4, 1, 2, 1, 2, 6, 7, 6, 7}
	require.Equal(t, 4, singleNumber2(nums))
}
