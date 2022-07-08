package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_intersection(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{5, 2, 2, 3}
	require.EqualValues(t, []int{2}, intersection(nums1, nums2))
}
