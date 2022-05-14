package leetcode

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth, rightDepth := 0, 0
	if root.Left != nil {
		leftDepth = minDepth(root.Left)
	}
	if root.Right != nil {
		rightDepth = minDepth(root.Right)
	}

	switch {
	case rightDepth == 0:
		return leftDepth + 1
	case leftDepth == 0:
		return rightDepth + 1
	case leftDepth <= rightDepth:
		return leftDepth + 1
	default:
		return rightDepth + 1
	}

}
