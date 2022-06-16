package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeElement(t *testing.T) {
	nums := []int{1, 2, 3, 3, 5, 7, 3, 11}
	require.Equal(t, 5, removeElement(nums, 3))
}
