package leetcode

func intersect(nums1, nums2 []int) []int {
	countMap1 := make(map[int][]int)
	countMap2 := make(map[int][]int)
	for _, n := range nums1 {
		if intslice, ok := countMap1[n]; ok {
			countMap1[n] = append(intslice, n)
		} else {
			countMap1[n] = []int{n}
		}
	}
	for _, n := range nums2 {
		if intslice, ok := countMap2[n]; ok {
			countMap2[n] = append(intslice, n)
		} else {
			countMap2[n] = []int{n}
		}
	}
	size1, size2 := len(countMap1), len(countMap2)
	result := make([]int, 0)
	switch {
	case size1 < size2:
		for key, intslice1 := range countMap1 {
			if intslice2, ok := countMap2[key]; ok {
				if len(intslice1) > len(intslice2) {
					result = append(result, intslice2...)
				} else {
					result = append(result, intslice1...)
				}
			}
		}
	default:
		for key, intslice2 := range countMap2 {
			if intslice1, ok := countMap1[key]; ok {
				if len(intslice1) > len(intslice2) {
					result = append(result, intslice2...)
				} else {
					result = append(result, intslice1...)
				}
			}
		}
	}
	return result
}
