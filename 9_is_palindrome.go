package leetcode

func isPalindromeInt(x int) bool {
	if x < 0 {
		return false
	}
	var tmp, target int
	tmp = x
	for tmp != 0 {
		target = target*10 + tmp%10
		tmp = tmp / 10
	}
	return x == target
}
