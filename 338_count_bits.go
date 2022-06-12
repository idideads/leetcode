package leetcode

import "fmt"

func countBits(n int) []int {
	if n < 0 {
		return nil
	}
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		s := fmt.Sprintf("%b", i)
		count := 0
		for i := range s {
			if s[i] == '1' {
				count++
			}
		}
		ans[i] = count
	}
	return ans
}
