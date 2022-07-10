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

func NisPowerOfX(n, x int) bool {
	if n < 1 || x < 1 {
		return false
	}
	for n > 1 && n%x == 0 {
		n /= x
	}
	return n == 1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Difference(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
