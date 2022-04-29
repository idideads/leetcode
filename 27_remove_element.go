package leetcode

func removeElement(nums []int, val int) int {
	k, lenght := 0, len(nums)
	for _, n := range nums {
		if n != val {
			k++
		}
	}
	i, j := 0, lenght-1
	for i < j {
		if nums[i] == val && nums[j] != val {
			nums[i], nums[j] = nums[j], nums[i]
		}
		if nums[i] != val {
			i++
		}
		if nums[j] == val {
			j--
		}
	}
	return k
}
