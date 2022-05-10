package leetcode

// ListNode is a singly-linked list
type ListNode struct {
	Val  int       `json:"val"`
	Next *ListNode `json:"next"`
}

// TreeNode is a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
