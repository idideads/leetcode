package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_deleteDuplicates(t *testing.T) {
	head := &ListNode{1, &ListNode{1, &ListNode{3, &ListNode{3, &ListNode{5, &ListNode{5, &ListNode{7, &ListNode{7, &ListNode{11, nil}}}}}}}}}
	head = deleteDuplicates(head)
	require.NotNil(t, head)
	require.EqualValues(t, []int{1, 3, 5, 7, 11}, ListToSlice(head))
}
