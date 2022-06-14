package leetcode

import (
	"fmt"
	"testing"
)

func Test_removeElements(t *testing.T) {
	head := &ListNode{1, &ListNode{2, &ListNode{6, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}}
	head = removeElements(head, 6)
	for head.Next != nil {
		fmt.Printf("%d->", head.Val)
		head = head.Next
	}
	fmt.Printf("%d\n", head.Val)
}
