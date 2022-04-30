package leetcode

/*
实现 strStr() 函数。
给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。

示例1：
输入：haystack = "hello", needle = "ll"
输出：2
示例2：
输入：haystack = "aaaaa", needle = "bba"
输出：-1
示例3：
输入：haystack = "", needle = ""
输出：0
示例4；
输入：haystack = "mississippi", needle = "issip"
输出：4
*/
func strStr(haystack string, needle string) int {
	i, haystackLength, needleLength := 0, len(haystack), len(needle)
	startPosition := -1
	if needleLength == 0 {
		return 0
	}
	for i < haystackLength {
		if haystack[i] == needle[0] && haystackLength >= i+needleLength && haystack[i:i+needleLength] == needle {
			startPosition = i
			break
		}
		i++
	}
	return startPosition
}
