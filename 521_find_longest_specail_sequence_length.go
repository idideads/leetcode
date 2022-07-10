package leetcode

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	aLength, bLength := len(a), len(b)
	return Max(aLength, bLength)
}
