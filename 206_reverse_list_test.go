package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverseList(t *testing.T) {
	ln4 := &ListNode{Val: 5}
	ln3 := &ListNode{Val: 4, Next: ln4}
	ln2 := &ListNode{Val: 3, Next: ln3}
	ln1 := &ListNode{Val: 2, Next: ln2}
	head := &ListNode{Val: 1, Next: ln1}
	head = reverseList(head)
	require.NotNil(t, head)
	require.EqualValues(t, []int{5, 4, 3, 2, 1}, ListToSlice(head))
}
