package leetcode

func romanToInt(s string) int {
	var dc map[string]int = map[string]int{"I":1, "V":5, "X":10, "L":50, "C":100, "D":500, "M":1000}
    var num, pernum, total int
	for i := 0; i < len(s); i++ {
		num = dc[s[len(s)-(i+1) : len(s)-i]]
		if num >= pernum {
			total += num
		} else {
			total -= num
		}

		pernum = num
	}
	return total
}
