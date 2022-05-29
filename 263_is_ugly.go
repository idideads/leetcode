package leetcode

func isUgly(n int) bool {
	if n < 1 {
		return false
	}
	for n%2 == 0 || n%3 == 0 || n%5 == 0 {
		switch {
		case n%2 == 0:
			n /= 2
		case n%3 == 0:
			n /= 3
		case n%5 == 0:
			n /= 5
		}
	}
	return n == 1
}
