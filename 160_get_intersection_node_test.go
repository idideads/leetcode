package leetcode

import (
	"fmt"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	ln2 := &ListNode{Val: 5, Next: nil}
	ln1 := &ListNode{Val: 4, Next: ln2}
	ln0 := &ListNode{Val: 8, Next: ln1}
	aln1 := &ListNode{Val: 1, Next: ln0}
	aln0 := &ListNode{Val: 4, Next: aln1}
	bln2 := &ListNode{Val: 1, Next: ln0}
	bln1 := &ListNode{Val: 6, Next: bln2}
	bln0 := &ListNode{Val: 5, Next: bln1}
	fmt.Println(getIntersectionNode(aln0, bln0).Val)
}
