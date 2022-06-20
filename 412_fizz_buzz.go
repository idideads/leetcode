package leetcode

import "fmt"

func fizzBuzz(n int) []string {
	answer := make([]string, n)
	for i := 0; i < n; i++ {
		switch {
		case (i+1)%3 == 0 && (i+1)%5 == 0:
			answer[i] = "FizzBuzz"
		case (i+1)%3 == 0:
			answer[i] = "Fizz"
		case (i+1)%5 == 0:
			answer[i] = "Buzz"
		default:
			answer[i] = fmt.Sprint(i + 1)
		}
	}
	return answer
}
