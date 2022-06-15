package leetcode

func firstUniqChar(s string) int {
	countMap := make(map[rune]int)
	for _, c := range s {
		countMap[c]++
	}
	for i, c := range s {
		if countMap[c] == 1 {
			return i
		}
	}
	return -1
}
