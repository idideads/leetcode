package leetcode

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	nodes := make([]int, 0)
	nodes = append(nodes, root.Val)

	length := len(root.Children)
	if length == 0 {
		return nodes
	}

	for i := range root.Children {
		nodeChildren := preorder(root.Children[i])
		nodes = append(nodes, nodeChildren...)
	}

	return nodes
}
