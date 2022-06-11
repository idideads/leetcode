package leetcode

func NisPowerOfX(n, x int) bool {
	if n < 1 || x < 1 {
		return false
	}
	for n > 1 && n%x == 0 {
		n /= x
	}
	return n == 1
}
