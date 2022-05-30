package leetcode

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	sr := make([]string, 0)
	startIndex := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] > 1 {
			if i == startIndex {
				sr = append(sr, fmt.Sprintf("%d", nums[i]))
			} else {
				sr = append(sr, fmt.Sprintf("%d->%d", nums[startIndex], nums[i]))
			}
			startIndex = i + 1
		}
	}
	if startIndex == len(nums)-1 {
		sr = append(sr, fmt.Sprintf("%d", nums[startIndex]))
	} else {
		sr = append(sr, fmt.Sprintf("%d->%d", nums[startIndex], nums[len(nums)-1]))
	}
	return sr
}
