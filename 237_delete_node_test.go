package leetcode

import (
	"fmt"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	ln3 := &ListNode{Val: 9, Next: nil}
	ln2 := &ListNode{Val: 1, Next: ln3}
	ln1 := &ListNode{Val: 5, Next: ln2}
	ln0 := &ListNode{Val: 4, Next: ln1}
	deleteNode(ln2)
	for ln0.Next != nil {
		fmt.Printf("%d->", ln0.Val)
		ln0 = ln0.Next
	}
	fmt.Printf("%d\n", ln0.Val)
}
