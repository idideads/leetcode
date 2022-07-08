package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_wordPattern(t *testing.T) {

	require.Equal(t, false, wordPattern("abba", "dog dog dog dog"))
}
