package leetcode

func isHappy(n int) bool {
	falseNumberMap := make(map[int]struct{})
	m := n
	for m != 1 {
		m = 0
		falseNumberMap[n] = struct{}{}
		for n > 0 {
			m += (n % 10) * (n % 10)
			n /= 10
		}
		if _, ok := falseNumberMap[m]; ok {
			return false
		}
		n = m
	}
	return true
}
