package leetcode

func containsDuplicate(nums []int) bool {
	numMap := make(map[int]struct{})
	for _, n := range nums {
		if _, ok := numMap[n]; ok {
			return true
		}
		numMap[n] = struct{}{}
	}
	return false
}
