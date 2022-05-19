package leetcode

import "math"

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return math.Abs(float64(leftDepth)-float64(rightDepth)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}
