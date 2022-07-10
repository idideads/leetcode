package leetcode

func constructRectangle(area int) []int {
	maxW := 1
	for maxW*maxW <= area {
		maxW++
	}
	for w := maxW; w > 0; w-- {
		if area%w == 0 {
			if area/w > w {
				return []int{area / w, w}
			} else {
				return []int{w, area / w}
			}
		}
	}
	return []int{area, 1}
}
