package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverseWords(t *testing.T) {
	require.Equal(t, "a", reverseWords("a"))
	require.Equal(t, "ba dc", reverseWords("ab cd"))
	require.Equal(t, "s'teL ekat edoCteeL tsetnoc", reverseWords("Let's take LeetCode contest"))
	require.Equal(t, "doG gniD", reverseWords("God Ding"))
}
