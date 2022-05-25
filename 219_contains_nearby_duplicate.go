package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	numMap := make(map[int]int)
	for i, n := range nums {
		if index, ok := numMap[n]; ok && i-index <= k {
			return true
		}
		numMap[n] = i
	}
	return false
}
