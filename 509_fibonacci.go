package leetcode

func fib(n int) int {
    if n < 2 {
		return n
	}
	elementNminus2, elementNminus1, elementN := 0, 1, 0
	for i := 2; i <= n; i++ {
		elementN = elementNminus1 + elementNminus2
		elementNminus1, elementNminus2 = elementN, elementNminus1
	}
	return elementN
}