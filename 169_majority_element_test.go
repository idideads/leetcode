package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_majorityElement(t *testing.T) {
	nums := []int{2, 2, 2, 11, 2, 2, 1, 3, 4, 2, 2, 2}
	require.Equal(t, 2, majorityElement(nums))
}
