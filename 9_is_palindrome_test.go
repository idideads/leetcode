package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPalindromeInt(t *testing.T) {
	require.Equal(t, true, isPalindromeInt(121))
}
