package leetcode

import (
	"fmt"
	"testing"
)

func Test_addBinary(t *testing.T) {
	a, b := "1111", "1111"
	c := addBinary(a, b)
	fmt.Println(a, b, c)
}

func Test_op(t *testing.T) {
	s := "1"
	b := []byte(s)
	fmt.Println(int(b[0]) - 48)
	c0, c1 := byte('0'), byte('1')
	fmt.Printf("c0:%d\t%b\t%c\nc1:%d\t%b\t%c\n", c0, c0, c0, c1, c1, c1)
	// fmt.Println("")
	// fmt.Printf("0 & 0 = %c\t%b\n", c0&c0, c0&c0)
	// fmt.Printf("0 & 1 = %c\t%b\n", c0&c1, c0&c1)
	// fmt.Printf("1 & 1 = %c\t%b\n", c1&c1, c1&c1)
	// fmt.Println("")
	// fmt.Printf("0 | 0 = %c\t%b\n", c0|c0, c0|c0)
	// fmt.Printf("0 | 1 = %c\t%b\n", c0|c1, c0|c1)
	// fmt.Printf("1 | 1 = %c\t%b\n", c1|c1, c1|c1)
	// fmt.Println("")
	// fmt.Printf("0 ^ 0 = %c\t%b\n", c0^c0, c0^c0)
	// fmt.Printf("0 ^ 1 = %c\t%b\n", c0^c1, c0^c1)
	// fmt.Printf("1 ^ 1 = %c\t%b\n", c1^c1, c1^c1)
	// fmt.Println("")
	// fmt.Printf("0 ^ 0 | 0 = %d\t%b\n", c0^c0|'0', c0^c0|'0')
	// fmt.Printf("0 ^ 1 | 0 = %d\t%b\n", c0^c1|'0', c0^c1|'0')
	// fmt.Printf("1 ^ 1 | 0 = %d\t%b\n", c1^c1|'0', c1^c1|'0')
}
