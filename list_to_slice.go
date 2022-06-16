package leetcode

func ListToSlice(head *ListNode) []int {
	if head == nil {
		return nil
	}
	listSlice := make([]int, 0)
	for head.Next != nil {
		listSlice = append(listSlice, head.Val)
		head = head.Next
	}
	listSlice = append(listSlice, head.Val)
	return listSlice
}
