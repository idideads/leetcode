package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findLUSlength(t *testing.T) {
	require.Equal(t, 3, findLUSlength("aba", "cdc"))
	require.Equal(t, 3, findLUSlength("aaa", "bbb"))
	require.Equal(t, -1, findLUSlength("aaa", "aaa"))
}
