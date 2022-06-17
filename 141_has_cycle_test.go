package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_hasCycle(t *testing.T) {
	listNode_5 := &ListNode{Val: 7, Next: nil}
	listNode_4 := &ListNode{Val: 5, Next: listNode_5}
	listNode_3 := &ListNode{Val: 3, Next: listNode_4}
	listNode_2 := &ListNode{Val: 2, Next: listNode_3}
	listNode_1 := &ListNode{Val: 1, Next: listNode_2}
	listNode_5.Next = listNode_2
	require.Equal(t, true, hasCycle(listNode_1))
}
