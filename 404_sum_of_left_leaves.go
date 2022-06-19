package leetcode

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 0
	}
	//关键是如何判断左叶子节点，遇到左叶子节点就累加
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	} else {
		return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
	}
}
