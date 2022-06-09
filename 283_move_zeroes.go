package leetcode

func moveZeroes(nums []int) {
	size := len(nums)
	i, j := 0, 0
	for i < size && j < size {
		if nums[i] == 0 && nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
		}
		if nums[i] != 0 {
			i++
		}
		j++
	}
}
