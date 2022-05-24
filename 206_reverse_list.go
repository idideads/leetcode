package leetcode

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	nodeSlice := make([]*ListNode, 0)
	for head.Next != nil {
		nodeSlice = append(nodeSlice, head)
		head = head.Next
	}
	nodeSlice = append(nodeSlice, head)
	length := len(nodeSlice)
	for i := length - 1; i > 0; i-- {
		nodeSlice[i].Next = nodeSlice[i-1]
	}
	nodeSlice[0].Next = nil
	return nodeSlice[length-1]
}
