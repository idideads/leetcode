package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListToSlice(t *testing.T) {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	require.NotNil(t, ListToSlice(head))
	require.EqualValues(t, []int{1, 2, 3, 4}, ListToSlice(head))
}

func TestNisPowerOfX(t *testing.T) {
	require.Equal(t, true, NisPowerOfX(16, 4))
	require.Equal(t, true, NisPowerOfX(8, 2))
	require.Equal(t, false, NisPowerOfX(20, 3))
}

func TestMin(t *testing.T) {
	require.Equal(t, 3, Min(3, 7))
	require.Equal(t, 8, Min(12, 8))
}

func TestDifference(t *testing.T) {
	require.Equal(t, 3, Difference(3, 6))
	require.Equal(t, 3, Difference(9, 6))
}

func TestMax(t *testing.T) {
	require.Equal(t, 6, Max(2, 6))
	require.Equal(t, 7, Max(7, 6))
}
