package leetcode

import "math"

func getMinimumDifference(root *TreeNode) int {
	min := math.MaxInt64
	nodes := inorderTraversal(root)
	for i := 0; i < len(nodes)-1; i++ {
		min = Min(min, Difference(nodes[i], nodes[i+1]))
	}
	return min
}
