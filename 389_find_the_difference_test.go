package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findTheDifference(t *testing.T) {
	got := findTheDifference("abcd", "abcde")
	var want byte = 'e'
	require.Equal(t, want, got)

	want = 'y'
	got = findTheDifference("", "y")
	require.Equal(t, want, got)
}
