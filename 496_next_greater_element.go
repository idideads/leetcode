package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
    lengthNums2 := len(nums2)
	nums2Map := make(map[int]int, lengthNums2)
	for i, n := range nums2 {
		nums2Map[n] = i
	}
	greaterElement := make([]int, len(nums1))
	for i, n := range nums1 {
		greaterElement[i] = -1
		index := nums2Map[n]
		for j := index; j < lengthNums2; j++ {
			if nums2[j] > n {
				greaterElement[i] = nums2[j]
				break
			}
		}
	}
	return greaterElement
}