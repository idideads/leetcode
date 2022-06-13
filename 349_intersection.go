package leetcode

func intersection(nums1 []int, nums2 []int) []int {
    intersectionMap := make(map[int]bool)
	for _, n := range nums1 {
		intersectionMap[n] = false
	}
	for _, n := range nums2 {
		if _, ok := intersectionMap[n]; ok {
			intersectionMap[n] = true
		}
	}
	result := make([]int, 0)
	for key, value := range intersectionMap {
		if value {
			result = append(result, key)
		}
	}
	return result
}