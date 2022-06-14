package leetcode

func guessNumber(n int) int {
	i, j := 1, n
	for i < j {
		h := (i + j) / 2
		if guess(h) == 0 {
			return h
		}
		if guess(h) == 1 {
			i = h + 1
		}
		if guess(h) == -1 {
			j = h
		}
	}
	return i
}

func guess(num int) int {
	pick := 6
	switch {
	case num == pick:
		return 0
	case num > pick:
		return -1
	default:
		return 1
	}
}
