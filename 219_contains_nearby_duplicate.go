package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	numMap := make(map[int]int)
	for i, n := range nums {
		index, ok := numMap[n]
		if ok {
			if i-index <= k {
				return true
			}
		}
		numMap[n] = i
	}
	return false
}
