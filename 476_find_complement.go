package leetcode

import (
	"fmt"
	"math"
)

func findComplement(num int) int {
	reverseNum := make([]byte, 0)
	for num > 0 {
		s := fmt.Sprint(num % 2)
		reverseNum = append(reverseNum, []byte(s)...)
		num /= 2
	}
	for i, j := 0, len(reverseNum)-1; i <= j; i, j = i+1, j-1 {
		if reverseNum[i] == '0' {
			reverseNum[i] = '1'
		} else {
			reverseNum[i] = '0'
		}

		if i != j {
			if reverseNum[j] == '0' {
				reverseNum[j] = '1'
			} else {
				reverseNum[j] = '0'
			}
		}

		reverseNum[i], reverseNum[j] = reverseNum[j], reverseNum[i]
	}
	for i, c := range reverseNum {
		power := len(reverseNum) - 1 - i
		num += int(c-'0') * int(math.Pow(2, float64(power)))
	}
	return num
}
