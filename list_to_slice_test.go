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
