package leetcode

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	nodeSlice := make([]*ListNode, 0)
	for head.Next != nil {
		if head.Val != val {
			nodeSlice = append(nodeSlice, head)
		}
		head = head.Next
	}
	if head.Val != val {
		nodeSlice = append(nodeSlice, head)
	}
	length := len(nodeSlice)
	for i := 0; i < length; i++ {
		if i == length-1 {
			nodeSlice[i].Next = nil
		} else {
			nodeSlice[i].Next = nodeSlice[i+1]
		}
	}
	if length == 0 {
		return nil
	}
	return nodeSlice[0]
}
