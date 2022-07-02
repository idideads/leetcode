package leetcode

import (
	"fmt"
	"sort"
)

func findRelativeRanks(score []int) []string {
	length := len(score)
	scoreCopy := make([]int, length)
	copy(scoreCopy, score)

	sort.Slice(scoreCopy, func(i, j int) bool {
		return scoreCopy[i] > scoreCopy[j]
	})

	// fmt.Printf("score:%v\n", score)
	// fmt.Printf("scoreCopy:%v\n", scoreCopy)

	rankMap := make(map[int]string, length)
	for i, scoreNum := range scoreCopy {
		switch {
		case i == 0:
			rankMap[scoreNum] = "Gold Medal"
		case i == 1:
			rankMap[scoreNum] = "Silver Medal"
		case i == 2:
			rankMap[scoreNum] = "Bronze Medal"
		default:
			rankMap[scoreNum] = fmt.Sprint(i + 1)
		}
	}

	rankSlice := make([]string, length)
	for i, scoreNum := range score {
		rankSlice[i] = rankMap[scoreNum]
	}
	return rankSlice
}
