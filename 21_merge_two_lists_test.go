package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mergeTwoLists(t *testing.T) {
	l1_1, l1_2, l1_3 := &ListNode{Val: 1}, &ListNode{Val: 2}, &ListNode{Val: 4}
	l1_1.Next, l1_2.Next = l1_2, l1_3
	l2_1, l2_2, l2_3 := &ListNode{Val: 1}, &ListNode{Val: 3}, &ListNode{Val: 4}
	l2_1.Next, l2_2.Next = l2_2, l2_3

	require.NotNil(t, mergeTwoLists(l1_1, l2_1))
	require.EqualValues(t, []int{1, 1, 2, 3, 4, 4}, ListToSlice(mergeTwoLists(l1_1, l2_1)))
}
