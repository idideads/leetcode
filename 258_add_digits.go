package leetcode

func addDigits(num int) int {
	n := num
	for n > 9 {
		n = 0
		for num > 0 {
			n += num % 10
			num /= 10
		}
		num = n
	}
	return n
}
