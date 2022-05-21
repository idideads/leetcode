package leetcode

/*
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
说明：本题中，我们将空字符串定义为有效的回文串。
示例1：
输入: "A man, a plan, a canal: Panama"
输出: true
解释："amanaplanacanalpanama" 是回文串
*/
func isPalindromeStr(s string) bool {
	// if len(s) == 0 || s == "" {
	// 	return true
	// }
	bytes := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] >= 48 && s[i] <= 57:
			bytes = append(bytes, s[i])
		case s[i] >= 65 && s[i] <= 90:
			bytes = append(bytes, s[i]+32)
		case s[i] >= 97 && s[i] <= 122:
			bytes = append(bytes, s[i])
		}
	}
	size := len(bytes)
	if size == 0 {
		return true
	}
	// fmt.Println(bytes)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		if bytes[i] != bytes[j] {
			return false
		}
	}
	return true
}
