package leetcode

func distributeCandies(candyType []int) int {
	n := len(candyType)
	candyMap := make(map[int]int)
	for _, candy := range candyType {
		candyMap[candy]++
	}
	return Min(len(candyMap), n/2)
}
