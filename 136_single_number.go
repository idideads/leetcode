package leetcode

import "sort"

/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
说明：你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
示例1：
输入: [4,1,2,1,2]
输出: 4
*/
func singleNumber(nums []int) int {
	l := len(nums)
	sigNum := nums[0]
	for i := 1; i < l; i++ {
		sigNum ^= nums[i]
	}
	return sigNum
}

func singleNumber2(nums []int) int {
	size := len(nums)
	if size == 1 {
		return nums[0]
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] <= nums[j]
	})
	if nums[0] != nums[1] {
		return nums[0]
	}
	if nums[size-1] != nums[size-2] {
		return nums[size-1]
	}
	for i := 1; i < size-1; i++ {
		if nums[i] != nums[i-1] && nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return 0
}
