package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findTheDifference(t *testing.T) {
	require.Equal(t, 'e', findTheDifference("abcd", "abcde"))
	require.Equal(t, 'y', findTheDifference("", "y"))
}
