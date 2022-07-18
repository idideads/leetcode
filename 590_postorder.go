package leetcode


func postorder(root *Node) []int {
	nodes := make([]int,0)
	if root==nil {
		return nodes
	}
	nodeStack := make([]*Node,1)
	nodeStack[0]=root
	for len(nodeStack)>0{
		node := nodeStack[len(nodeStack)-1]
		for i := len(node.Children)-1; i >= 0; i-- {
			nodeStack = append(nodeStack, node.Children[i])
		}
		if len(node.Children)==0 {
			nodeStack=nodeStack[:len(nodeStack)-1]
		}
	}
	return nodes
}