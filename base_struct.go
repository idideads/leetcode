package leetcode

// ListNode is a singly-linked list
type ListNode struct {
	Val  int       `json:"val"`
	Next *ListNode `json:"next"`
}

// TreeNode is a binary tree
type TreeNode struct {
	Val   int       `json:"val"`
	Left  *TreeNode `json:"left"`
	Right *TreeNode `json:"right"`
}

// Node is N fork-tree
type Node struct {
	Val      int     `json:"val"`
	Children []*Node `json:"children"`
}
