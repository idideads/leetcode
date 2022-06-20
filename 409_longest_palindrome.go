package leetcode

func longestPalindrome(s string) int {
	countMap := make(map[rune]int)
	for _, sRune := range s {
		countMap[sRune]++
	}
	oddNum, evenNum := 0, 0
	for _, v := range countMap {
		if v%2 == 0 {
			evenNum += v
		} else {
			if v > 2 {
				evenNum += v - 1
			}
			oddNum = 1
		}
	}
	return evenNum + oddNum
}
