package leetcode

func climbStairs(n int) int {
    if n < 3 {
		return n
	}
	stairsNminus2, stairsNminus1, stairsN := 1, 2, 0
	for i := 3; i <= n; i++ {
		stairsN = stairsNminus1 + stairsNminus2
		stairsNminus2, stairsNminus1 = stairsNminus1, stairsN
	}
	return stairsN
}