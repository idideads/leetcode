package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_intersect(t *testing.T) {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	require.EqualValues(t, []int{9, 4}, intersect(nums1, nums2))
}
