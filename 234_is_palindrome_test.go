package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPalindromeList(t *testing.T) {
	head := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}
	require.Equal(t, true, isPalindromeList(head))
	head = &ListNode{1, &ListNode{3, &ListNode{5, nil}}}
	require.Equal(t, false, isPalindromeList(head))
}
