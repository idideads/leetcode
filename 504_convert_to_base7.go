package leetcode

import "fmt"

func convertToBase7(num int) string {
	answer, isNegative := "", false
	if num < 0 {
		isNegative = true
		num *= -1
	}
	for num/7 != 0 {
		answer = fmt.Sprint(num%7) + answer
		num /= 7
	}
	if num >= 0 && num < 7 {
		answer = fmt.Sprint(num) + answer
	}
	if isNegative {
		answer = "-" + answer
	}
	return answer
}
