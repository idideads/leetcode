package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestPalindrome(t *testing.T) {
	require.Equal(t, 1, longestPalindrome("a"))
	require.Equal(t, 1, longestPalindrome("abcd"))
	require.Equal(t, 3, longestPalindrome("bbb"))
	require.Equal(t, 7, longestPalindrome("abccccdd"))
}
