package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lengthOfLastWord(t *testing.T) {
	require.Equal(t, 5, lengthOfLastWord("hello  world  "))
	require.Equal(t, 4, lengthOfLastWord("   fly me   to   the moon  "))
	require.Equal(t, 6, lengthOfLastWord("luffy is still joyboy"))
}
