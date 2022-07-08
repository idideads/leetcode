package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPowerOfThree(t *testing.T) {
	require.Equal(t, true, isPowerOfThree(27))
	require.Equal(t, false, isPowerOfThree(12))
	require.Equal(t, true, isPowerOfThree(1))
}
