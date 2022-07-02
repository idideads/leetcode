package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_detectCapitalUse(t *testing.T) {
	require.Equal(t,true,detectCapitalUse("a"))
	require.Equal(t,true,detectCapitalUse("A"))
	require.Equal(t,true,detectCapitalUse("ABC"))
	require.Equal(t,true,detectCapitalUse("Abc"))
	require.Equal(t,true,detectCapitalUse("abc"))
	require.Equal(t,false,detectCapitalUse("aBc"))
	require.Equal(t,false,detectCapitalUse("abC"))
	require.Equal(t,false,detectCapitalUse("AbcD"))
}
