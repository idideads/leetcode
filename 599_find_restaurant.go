package leetcode

import "math"

func findRestaurant(list1, list2 []string) []string {
	map1 := make(map[string]int)
	for i, s := range list1 {
		map1[s] = i
	}

	res := []string{}
	totalIndex := math.MaxInt64

	for i, s := range list2 {
		if _, ok := map1[s]; ok {
			if i+map1[s] == totalIndex {
				res = append(res, s)
			}

			if i+map1[s] < totalIndex {
				res = []string{}
				res = append(res, s)
				totalIndex = i + map1[s]
			}

		}
	}

	return res
}
