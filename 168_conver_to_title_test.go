package leetcode

import (
	"fmt"
	"testing"
)

func Test_convertToTitle(t *testing.T) {
	columnNumber := 1
	fmt.Println(convertToTitle(columnNumber))
	columnNumber = 28
	fmt.Println(convertToTitle(columnNumber))
	columnNumber = 52
	fmt.Println(convertToTitle(columnNumber))
	columnNumber = 701
	fmt.Println(convertToTitle(columnNumber))
	columnNumber = 2147483647
	fmt.Println(convertToTitle(columnNumber))
}
