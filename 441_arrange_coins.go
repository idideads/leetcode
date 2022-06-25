package leetcode

func arrangeCoins(n int) int {
	k := 1
	for (1+k)*k <= 2*n {
		k++
	}
	return k - 1
}
