package leetcode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	stubMap := make(map[*ListNode]struct{})
	for headA.Next != nil {
		stubMap[headA] = struct{}{}
		headA = headA.Next
	}
	stubMap[headA] = struct{}{}

	for headB.Next != nil {
		if _, ok := stubMap[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	if _, ok := stubMap[headB]; ok {
		return headB
	}
	return nil
}
