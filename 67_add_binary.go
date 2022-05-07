package leetcode

/*
给你两个二进制字符串，返回它们的和（用二进制表示）。
输入为 非空 字符串且只包含数字 1 和 0。
示例1：
输入: a = "11", b = "1"
输出: "100"
示例2：
输入: a = "1010", b = "1011"
输出: "10101"
*/
func addBinary(a string, b string) string {
	aLength, bLength := len(a), len(b)
	var abytes, bbytes, rbytes []byte
	longestSize := 0
	if aLength >= bLength {
		longestSize = aLength + 1
	} else {
		longestSize = bLength + 1
	}
	abytes, bbytes, rbytes = make([]byte, longestSize-aLength, longestSize), make([]byte, longestSize-bLength, longestSize), make([]byte, longestSize)
	for i, _ := range abytes {
		abytes[i] = '0'
	}
	for i, _ := range bbytes {
		bbytes[i] = '0'
	}
	abytes, bbytes = append(abytes, []byte(a)...), append(bbytes, []byte(b)...)
	var c int
	for i := longestSize - 1; i >= 0; i-- {
		a, b := int(abytes[i])-48, int(bbytes[i])-48
		if a+b+c > 1 {
			rbytes[i], c = byte(a+b+c-2)+48, 1
		} else {
			rbytes[i], c = byte(a+b+c)+48, 0
		}
	}
	if rbytes[0] == '0' {
		return string(rbytes[1:])
	}
	return string(rbytes)
}
