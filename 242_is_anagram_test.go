package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isAnagram(t *testing.T) {
	require.Equal(t, true, isAnagram("qwerty", "ytrewq"))
	require.Equal(t, false, isAnagram("bat", "cat"))
	require.Equal(t, false, isAnagram("abc", "dbca"))
}

func Test_isAnagramByMap(t *testing.T) {
	require.Equal(t, true, isAnagramByMap("qwerty", "ytrewq"))
}
