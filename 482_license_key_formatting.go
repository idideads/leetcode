package leetcode

func licenseKeyFormatting(s string, k int) string {
	if k < 1 || len(s) == 0 {
		return ""
	}
	bytes, count := make([]byte, 0), 0
	for len(s) > 0 {
		i := len(s) - 1
		if s[i] == '-' {
			s = s[:i]
			continue
		}
		if s[i] >= 'a' && s[i] <= 'z' {
			bytes = append(bytes, s[i]-'a'+'A')
		} else {
			bytes = append(bytes, s[i])
		}
		count++
		if count > 0 && count%k == 0 {
			bytes = append(bytes, '-')
		}
		s = s[:i]
	}
	length := len(bytes)
	if length == 0 {
		return ""
	}
	if bytes[length-1] == '-' {
		bytes = bytes[:length-1]
	}
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}

func licenseKeyFormattingReverse(s string, k int) string {
	if k < 1 || len(s) == 0 {
		return ""
	}
	bytes, count := make([]byte, 0), 0
	for len(s) > 0 {
		i := len(s) - 1
		if s[i] == '-' {
			s = s[:i]
			continue
		}
		if s[i] >= 'a' && s[i] <= 'z' {
			bytes = append([]byte{s[i] - 32}, bytes...) //不要从切片的头部插入数据，元素都要往后挪效率太低
		} else {
			bytes = append([]byte{s[i]}, bytes...) //不要从切片的头部插入数据，元素都要往后挪效率太低
		}
		count++
		if count > 0 && count%k == 0 {
			bytes = append([]byte{'-'}, bytes...) //不要从切片的头部插入数据，元素都要往后挪效率太低
		}
		s = s[:i]
	}
	if len(bytes) == 0 {
		return ""
	}
	if bytes[0] == '-' {
		return string(bytes[1:])
	}
	return string(bytes)
}
