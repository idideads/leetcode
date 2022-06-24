package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_addStrings(t *testing.T) {
	require.Equal(t, "0", addStrings("0", "0"))
	require.Equal(t, "134", addStrings("11", "123"))
	require.Equal(t, "529", addStrings("456", "73"))
	require.Equal(t, "1033", addStrings("956", "77"))
}
