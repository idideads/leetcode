package leetcode

/*
给定一个二叉树的根节点 root ，返回 它的 中序 遍历： 左节点->根节点->右节点
*/
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	slice := make([]int, 0)
	if root.Left != nil {
		slice = append(slice, inorderTraversal(root.Left)...)
	}
	slice = append(slice, root.Val)
	if root.Right != nil {
		slice = append(slice, inorderTraversal(root.Right)...)
	}
	return slice
}
