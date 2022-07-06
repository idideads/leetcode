package leetcode

func distributeCandies(candyType []int) int {
    n := len(candyType)
	candyMap := make(map[int]int)
	for _, candy := range candyType {
		candyMap[candy]++
	}
	if len(candyMap) < n/2 {
		return len(candyMap)
	} else {
		return n / 2
	}
}