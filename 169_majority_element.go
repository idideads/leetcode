package leetcode

func majorityElement(nums []int) int {
	countMap := make(map[int]int, len(nums))
	for _, n := range nums {
		countMap[n]++
	}
	maxCount, maxKey := 0, 0
	for k, c := range countMap {
		if c > maxCount {
			maxCount = c
			maxKey = k
		}
	}
	return maxKey
}
