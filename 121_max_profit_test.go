package leetcode

import (
	"fmt"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	input := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(input))

	input = []int{7, 6, 5, 4, 3}
	fmt.Println(maxProfit(input))

	input = []int{2,4,1}
	fmt.Println(maxProfit(input))
}
