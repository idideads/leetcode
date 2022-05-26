package leetcode

func deleteNode(node *ListNode) {
	if node == nil {
		return
	}

	for node.Next.Next != nil {
		node.Val = node.Next.Val
		node = node.Next
	}

	node.Val = node.Next.Val
	node.Next = nil

}
