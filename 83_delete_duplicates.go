package leetcode

/*
给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
示例1：
输入：head = [1,1,2]
输出：[1,2]
示例2：
输入：head = [1,1,2,3,3]
输出：[1,2,3]
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	slice := make([]*ListNode, 0)
	for head.Next != nil {
		if head.Val != head.Next.Val {
			slice = append(slice, head)
		}
		head = head.Next
	}
	if len(slice) == 0 {
		slice = append(slice, head)
		return slice[0]
	}
	if slice[len(slice)-1].Val != head.Val {
		slice = append(slice, head)
	}
	for i, node := range slice {
		if i < len(slice)-1 {
			node.Next = slice[i+1]
		}
	}
	slice[len(slice)-1].Next = nil
	return slice[0]
}
