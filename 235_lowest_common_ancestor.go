package leetcode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	switch {
	case root == nil || p == nil || q == nil:
		return nil
	// case p.Val >= root.Val && root.Val >= q.Val:
	// 	return root
	// case q.Val >= root.Val && root.Val >= p.Val:
	// 	return root
	case p.Val < root.Val && q.Val < root.Val:
		return lowestCommonAncestor(root.Left, p, q)
	case p.Val > root.Val && q.Val > root.Val:
		return lowestCommonAncestor(root.Right, p, q)
	default:
		return root
	}
}
