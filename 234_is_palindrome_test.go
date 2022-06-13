package leetcode

import (
	"fmt"
	"testing"
)

func Test_isPalindromeList(t *testing.T) {
	ln3 := &ListNode{Val: 1}
	ln2 := &ListNode{Val: 2, Next: ln3}
	ln1 := &ListNode{Val: 2, Next: ln2}
	head := &ListNode{Val: 1, Next: ln1}
	fmt.Println(isPalindromeList(head))
}
