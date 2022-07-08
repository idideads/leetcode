package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canConstruct(t *testing.T) {
	ransomNote, magazine := "aaa", "aabbs"
	require.Equal(t, false, canConstruct(ransomNote, magazine))
}
