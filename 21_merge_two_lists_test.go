package leetcode

import (
	"fmt"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	l1_1, l1_2, l1_3 := &ListNode{Val: 1}, &ListNode{Val: 2}, &ListNode{Val: 4}
	l1_1.Next, l1_2.Next = l1_2, l1_3
	l2_1, l2_2, l2_3 := &ListNode{Val: 1}, &ListNode{Val: 3}, &ListNode{Val: 4}
	l2_1.Next, l2_2.Next = l2_2, l2_3

	var l3,l4 *ListNode
	// l3 = &ListNode{Val: 0}
	fmt.Printf("%p\t%v\n",l3,l3)
	fmt.Println("------------------------------------------------------")

	l := mergeTwoLists(l1_1, l2_1)
	for l != nil {
		fmt.Printf("%p\t%d\t%p\n", l, l.Val, l.Next)
		l = l.Next
	}
	fmt.Println("------------------------------------------------------")
	l1 := mergeTwoLists(l3, l4)
	for l1 != nil {
		fmt.Printf("%p\t%d\t%p\n", l1, l1.Val, l1.Next)
		l1 = l1.Next
	}
}
