package leetcode

/*
因为给定的是二叉搜索树，所以左边节点<=根节点<=右边节点
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	switch {
	case root == nil || p == nil || q == nil:
		return nil
	// case p.Val >= root.Val && root.Val >= q.Val:
	// 	return root
	// case q.Val >= root.Val && root.Val >= p.Val:
	// 	return root
	case p.Val < root.Val && q.Val < root.Val:
		//p、q都在根节点左边的，往左子树遍历
		return lowestCommonAncestor(root.Left, p, q)
	case p.Val > root.Val && q.Val > root.Val:
		//p、q都在根节点右边的，往右子树遍历
		return lowestCommonAncestor(root.Right, p, q)
	default:
		//p、q分别位于根节点左右两侧的直接返回根节点
		return root
	}
}
