package leetcode

func isSymmetric(root *TreeNode) bool {
	switch {
	case root == nil:
		return true
	case root.Left == nil && root.Right == nil:
		return true
	case root.Left == nil || root.Right == nil:
		return false
	case root.Left.Val != root.Right.Val:
		return false
	case root.Left.Val == root.Right.Val:
		return isMirror(root.Left, root.Right)
	}
	return false
}

func isMirror(left, right *TreeNode) bool {
	switch {
	case left == nil && right == nil:
		return true
	case left == nil || right == nil:
		return false
	case left.Val != right.Val:
		return false
	case left.Val == right.Val:
		return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
	}
	return false
}
