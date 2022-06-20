package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_containsNearbyDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	k := 3
	require.Equal(t, true, containsNearbyDuplicate(nums, k))
}
