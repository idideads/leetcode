package leetcode

func searchInsert(nums []int, target int) int {
	insertPosition := len(nums)
	for i, n := range nums {
		if n >= target {
			insertPosition = i
			break
		}
	}
	return insertPosition
}
