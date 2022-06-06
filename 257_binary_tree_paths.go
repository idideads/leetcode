package leetcode

import "fmt"

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []string{fmt.Sprint(root.Val)}
	}
	leftPaht, rightPath := make([]string, 0), make([]string, 0)
	if root.Left != nil {
		leftPaht = binaryTreePaths(root.Left)
	}
	if root.Right != nil {
		rightPath = binaryTreePaths(root.Right)
	}
	for i := range leftPaht {
		leftPaht[i] = fmt.Sprintf("%d->", root.Val) + leftPaht[i]
	}
	for i := range rightPath {
		rightPath[i] = fmt.Sprintf("%d->", root.Val) + rightPath[i]
	}
	return append(leftPaht, rightPath...)
}
