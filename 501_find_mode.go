package leetcode

func findMode(root *TreeNode) []int {
	modeMap := perorderTravelMap(root)
	maxMode := 0
	for _, v := range modeMap {
		if v > maxMode {
			maxMode = v
		}
	}
	modes := make([]int, 0)

	for k, v := range modeMap {
		if v == maxMode {
			modes = append(modes, k)
		}
	}
	return modes
}

func perorderTravelMap(root *TreeNode) map[int]int {
	if root == nil {
		return nil
	}
	unionMap := make(map[int]int)
	unionMap[root.Val]++
	leftMap, rightMap := perorderTravelMap(root.Left), perorderTravelMap(root.Right)
	for k, v := range leftMap {
		unionMap[k] += v
	}
	for k, v := range rightMap {
		unionMap[k] += v
	}
	return unionMap
}
