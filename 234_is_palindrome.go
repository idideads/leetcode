package leetcode

func isPalindromeList(head *ListNode) bool {
	if head == nil {
		return false
	}
	nodeSlice := make([]int, 0)
	for head.Next != nil {
		nodeSlice = append(nodeSlice, head.Val)
		head = head.Next
	}
	nodeSlice = append(nodeSlice, head.Val)
	size := len(nodeSlice)
	for i, j := 0, size-1; i <= j; i, j = i+1, j-1 {
		if nodeSlice[i] != nodeSlice[j] {
			return false
		}
	}
	return true
}
