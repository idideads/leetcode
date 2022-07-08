package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPowerOfFour(t *testing.T) {
	require.Equal(t, true, isPowerOfFour(16))
	require.Equal(t, true, isPowerOfFour(1))
}
