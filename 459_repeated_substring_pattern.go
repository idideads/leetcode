package leetcode

func repeatedSubstringPattern(s string) bool {
	length := len(s)
	if length < 2 {
		return false
	}
	for i := 1; i < length; i++ {
		if length%i == 0 {
			result := true
			for j := 0; j <= length-2*i; j = j + i {
				if s[j:j+i] != s[j+i:j+2*i] {
					result = false
					break
				}
			}
			if result {
				return result
			}
		}
	}
	return false
}
