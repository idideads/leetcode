package leetcode

func isValid(s string) bool {
    //括号对应的int32的值：`(`=40, `)`=41, `[`=91, `]`=93, `{`=123, `}`=125
	bMap := map[int32]int32{40:41, 91:93, 123:125}
	l := len(s)
	if l == 0 || l%2 != 0 {
		return false //空串或奇数个括号直接返回false
	}
	var stack []int32 //切片栈
	for _, v := range s {
		stackLen := len(stack)
		if stackLen > 0 && v == bMap[stack[stackLen-1]] {
			stack = stack[0 : len(stack)-1] //右括号匹配时，弹栈
		} else {
			stack = append(stack, v) //否则，压栈
		}
	}
	return len(stack)==0 //最后判断栈的元素个数，为0表示括号匹配成功
}
