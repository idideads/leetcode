package leetcode

func maxDepthTree(root *Node) int {
	if root == nil {
		return 0
	}

	if root.Children == nil {
		return 1
	}

	depth := 0
	for i := 0; i < len(root.Children); i++ {
		depth = Max(depth, maxDepthTree(root.Children[i]))
	}
	return depth + 1
}
