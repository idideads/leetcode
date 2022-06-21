package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_deleteNode(t *testing.T) {
	ln3 := &ListNode{Val: 9, Next: nil}
	ln2 := &ListNode{Val: 1, Next: ln3}
	ln1 := &ListNode{Val: 5, Next: ln2}
	ln0 := &ListNode{Val: 4, Next: ln1}

	deleteNode(ln2)
	require.EqualValues(t, []int{4, 5, 9}, ListToSlice(ln0))
}
