package leetcode

func findDisappearedNumbers(nums []int) []int {
	length := len(nums)
	numMap := make(map[int]struct{}, length)
	for _, n := range nums {
		numMap[n] = struct{}{}
	}
	dispNums := make([]int, 0)
	for i := 1; i <= length; i++ {
		if _, ok := numMap[i]; !ok {
			dispNums = append(dispNums, i)
		}
	}
	return dispNums
}
