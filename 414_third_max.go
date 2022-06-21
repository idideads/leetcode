package leetcode

import "sort"

func thirdMax(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	n := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			n++
		}
		if n == 2 {
			return nums[i+1]
		}
	}
	return nums[0]
}
