package leetcode

var maxDiameter int

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter = 0
	var diameterOfSubTree func(t *TreeNode) int
	diameterOfSubTree = func(t *TreeNode) int {
		if t != nil && t.Left == nil && t.Right == nil {
			return 1
		}
		leftDiameter, rightDiameter := 0, 0
		if t.Left != nil {
			leftDiameter = diameterOfSubTree(t.Left)
		}
		if t.Right != nil {
			rightDiameter = diameterOfSubTree(t.Right)
		}
		maxDiameter = Max(maxDiameter, leftDiameter+rightDiameter)
		return Max(leftDiameter, rightDiameter) + 1
	}
	diameterOfSubTree(root)
	return maxDiameter
}
