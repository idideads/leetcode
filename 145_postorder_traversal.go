package leetcode

/*
二叉树的前序遍历：左节点->右节点->根节点
*/
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	traversal := make([]int, 0)
	if root.Left != nil {
		traversal = append(traversal, postorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		traversal = append(traversal, postorderTraversal(root.Right)...)
	}
	traversal = append(traversal, root.Val)
	return traversal
}
