package leetcode

import "strings"

func reverseWords(s string) string {
	bytes := []byte(s)
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	stringSlice := strings.Split(string(bytes), " ")
	for i, j := 0, len(stringSlice)-1; i < j; i, j = i+1, j-1 {
		stringSlice[i], stringSlice[j] = stringSlice[j], stringSlice[i]
	}
	rs := ""
	for i := 0; i < len(stringSlice); i++ {
		if i == len(stringSlice)-1 {
			rs += stringSlice[i]
		} else {
			rs += (stringSlice[i] + " ")
		}
	}
	return rs
}
