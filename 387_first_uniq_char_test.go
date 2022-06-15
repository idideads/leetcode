package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_firstUniqChar(t *testing.T) {
	require.Equal(t, firstUniqChar("leetcode"), 0)
	require.Equal(t, firstUniqChar("loveleetcode"), 2)
	require.Equal(t, firstUniqChar("aaaddd"), -1)
}
