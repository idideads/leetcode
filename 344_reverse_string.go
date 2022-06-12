package leetcode

func reverseString(s []byte) {
	if len(s) == 0 {
		return
	}
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
