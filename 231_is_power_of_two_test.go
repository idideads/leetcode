package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPowerOfTwo(t *testing.T) {
	require.Equal(t, false, isPowerOfTwo(6))
	require.Equal(t, true, isPowerOfTwo(8))
}
