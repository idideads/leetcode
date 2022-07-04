package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_fib(t *testing.T) {
	require.Equal(t, 1, fib(2))
	require.Equal(t, 2, fib(3))
	require.Equal(t, 3, fib(4))
	require.Equal(t, 5, fib(5))
	require.Equal(t, 832040, fib(30))
}
