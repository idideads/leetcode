package leetcode

func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int, len(nums))
	for i, num := range nums {
		if v, ok := numMap[target-num]; ok {
			return []int{v, i}
		}
		numMap[num] = i
	}
	return nil
}
