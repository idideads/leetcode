package leetcode


func postorder(root *Node) []int {
	travels := make([]int, 0)
	if root == nil {
		return travels
	}
	nodeStack := make([]*Node, 1)
	nodeStack[0] = root
	for len(nodeStack) > 0 {
		node := nodeStack[len(nodeStack)-1]
		for i := len(node.Children) - 1; i >= 0; i-- {
			nodeStack = append(nodeStack, node.Children[i])
		}
		if len(node.Children) == 0 {
			nodeStack = nodeStack[:len(nodeStack)-1]
			travels = append(travels, node.Val)
		} else {
			node.Children = nil
		}
	}
	return travels
}