package leetcode

func findLHS(nums []int) int {
	numsMap := make(map[int]int)
	for _, n := range nums {
		numsMap[n]++
	}

	LHS := 0
	for _, n := range nums {
		if value, ok := numsMap[n+1]; ok {
			LHS = Max(LHS, numsMap[n]+value)
		}
	}

	return LHS
}
