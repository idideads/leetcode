package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isSubsequence(t *testing.T) {
	require.Equal(t, true, isSubsequence("abc", "ahbgdc"))
	require.Equal(t, false, isSubsequence("axc", "ahbgdc"))
	require.Equal(t, false, isSubsequence("aaaa", "bbaa"))
}
