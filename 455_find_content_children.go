package leetcode

import "sort"

func findContentChildren(g, s []int) int {
	if len(s) == 0 {
		return 0
	}
	if !sort.IntsAreSorted(g) {
		sort.Slice(g, func(i, j int) bool {
			return g[i] < g[j]
		})
	}
	if !sort.IntsAreSorted(s) {
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
	}
	n := 0
	for len(s) > 0 && len(g) > 0 {
		if s[0] >= g[0] {
			n++
			g = g[1:]
		}
		s = s[1:]
	}

	return n
}
