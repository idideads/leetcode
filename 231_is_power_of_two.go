package leetcode

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}
	for n > 1 && n%2 == 0 {
		n >>= 1
	}
	return n == 1
}
