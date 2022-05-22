package leetcode

/*
给你一个链表的头节点 head ，判断链表中是否有环。
*/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	visted := make(map[*ListNode]int, 0)
	visted[head] = head.Val
	for head.Next != nil {
		head = head.Next
		if _, ok := visted[head]; ok {
			return true
		} else {
			visted[head] = head.Val
		}
	}
	return false
}
