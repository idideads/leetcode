package leetcode

/**
*	将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*	输入：l1 = [1,2,4], l2 = [1,3,4]		输出：[1,1,2,3,4,4]
*	输入：l1 = [], l2 = []					输出：[]
*	输入：l1 = [], l2 = [0]					输出：[0]
 */



func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 用一个切片保存遍历的顺序
	mergeSlice := make([]*ListNode, 0)
	// 输入：l1 = [], l2 = []					输出：[]
	if list1 == nil && list2 == nil {
		return nil
	}

	// 当其中一个升序链表未遍历到尾时重复步骤
	for list1 != nil || list2 != nil {
		// 其中一个空（到尾）则把另一个放入切片
		if list2 == nil && list1 != nil {
			mergeSlice = append(mergeSlice, list1)
			list1 = list1.Next
		}

		if list1 == nil && list2 != nil {
			mergeSlice = append(mergeSlice, list2)
			list2 = list2.Next
		}

		// 两个都不为空则需要比较Val的值，值较小的放入切片同时链表往下走
		if list1 != nil && list2 != nil {
			if list1.Val <= list2.Val {
				mergeSlice = append(mergeSlice, list1)
				list1 = list1.Next
			} else {
				mergeSlice = append(mergeSlice, list2)
				list2 = list2.Next
			}
		}
	}

	// 最后遍历切片，把Next的指针理顺
	for i := 0; i < len(mergeSlice); i++ {
		if i == len(mergeSlice)-1 {
			mergeSlice[i].Next = nil
		} else {
			mergeSlice[i].Next = mergeSlice[i+1]
		}
	}

	return mergeSlice[0]
}
