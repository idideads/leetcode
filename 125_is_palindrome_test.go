package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPalindromeStr(t *testing.T) {
	require.Equal(t, true, isPalindromeStr("A man, a plan, a canal: Panama"))
	require.Equal(t, false, isPalindromeStr("apple"))
}
