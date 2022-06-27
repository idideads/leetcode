package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_repeatedSubstringPattern(t *testing.T) {
	require.Equal(t, false, repeatedSubstringPattern(""))
	require.Equal(t, false, repeatedSubstringPattern("a"))
	require.Equal(t, true, repeatedSubstringPattern("aa"))
	require.Equal(t, true, repeatedSubstringPattern("aaa"))
	require.Equal(t, false, repeatedSubstringPattern("aab"))
	require.Equal(t, true, repeatedSubstringPattern("abcabcabcabcabc"))
	require.Equal(t, true, repeatedSubstringPattern("abaababaab"))
}
