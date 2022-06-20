package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_fizzBuzz(t *testing.T) {
	answer := fizzBuzz(15)
	require.NotNil(t, answer)
	require.EqualValues(t, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}, answer)
}
