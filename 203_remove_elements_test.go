package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeElements(t *testing.T) {
	head := &ListNode{1, &ListNode{2, &ListNode{6, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}}
	head = removeElements(head, 6)
	require.NotNil(t, head)
	require.EqualValues(t, []int{1, 2, 3, 4, 5}, ListToSlice(head))
}
