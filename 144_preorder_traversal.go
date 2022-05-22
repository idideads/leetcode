package leetcode

/*
二叉树的前序遍历：根节点—>左节点->右节点
*/
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	traversal := make([]int, 0)
	traversal = append(traversal, root.Val)
	if root.Left != nil {
		traversal = append(traversal, preorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		traversal = append(traversal, preorderTraversal(root.Right)...)
	}
	return traversal
}