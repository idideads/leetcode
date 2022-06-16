package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isIsomorphic(t *testing.T) {
	require.Equal(t, true, isIsomorphic("egg", "add"))
	require.Equal(t, false, isIsomorphic("foo", "bar"))
	require.Equal(t, true, isIsomorphic("paper", "title"))
}
