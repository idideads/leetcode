package leetcode

import "strings"

func wordPattern(pattern string, s string) bool {
	slice := strings.Split(s, " ")
	size, length := len(pattern), len(slice)
	if size != length {
		return false
	}
	patternMap := make(map[byte]string, size)
	sliceMap := make(map[string]byte, length)
	for i := 0; i < size; i++ {
		patternStr, ok := patternMap[pattern[i]]
		if ok && patternStr != slice[i] {
			return false
		}
		patternMap[pattern[i]] = slice[i]

		sliceByte, ok := sliceMap[slice[i]]
		if ok && sliceByte != pattern[i] {
			return false
		}
		sliceMap[slice[i]] = pattern[i]

	}

	return true
}
