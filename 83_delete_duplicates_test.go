package leetcode

import (
	"fmt"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	listNode_5 := &ListNode{Val: 7, Next: nil}
	listNode_4 := &ListNode{Val: 3, Next: listNode_5}
	listNode_3 := &ListNode{Val: 3, Next: listNode_4}
	listNode_2 := &ListNode{Val: 1, Next: listNode_3}
	listNode_1 := &ListNode{Val: 1, Next: listNode_2}
	head := deleteDuplicates(listNode_1)
	for head.Next != nil {
		fmt.Printf("addr:%p\tval:%d\tnext:%p\n", head, head.Val, head.Next)
		head = head.Next
	}
	fmt.Printf("addr:%p\tval:%d\tnext:%p\n", head, head.Val, head.Next)
}
