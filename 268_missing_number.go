package leetcode

func missingNumber(nums []int) int {
	size := len(nums)
	numsMap := make(map[int]struct{}, size)
	for _, n := range nums {
		numsMap[n] = struct{}{}
	}
	for i := 0; i < size+1; i++ {
		if _, ok := numsMap[i]; !ok {
			return i
		}
	}
	return 0
}
