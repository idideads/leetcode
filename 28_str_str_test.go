package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_strStr(t *testing.T) {
	require.Equal(t, 2, strStr("hello", "ll"))
	require.Equal(t, -1, strStr("aaaaa", "bba"))
}
