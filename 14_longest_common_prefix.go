package leetcode

import "sort"

func longestCommonPrefix(strs []string) string {
    sLen := len(strs)
	var cPrefix string
	if sLen == 0 {
		return cPrefix
	}
	if sLen == 1 {
		return strs[0]
	}

	sort.Strings(strs)
	var minLen int
	if len(strs[0]) >= len(strs[sLen-1]) {
		minLen = len(strs[sLen-1])
	} else {
		minLen = len(strs[0])
	}
	for i := 0; i < minLen; i++ {
		if strs[0][i] == strs[sLen-1][i] {
			cPrefix += string(strs[0][i])
		} else {
			return cPrefix
		}
	}

	return cPrefix
}

