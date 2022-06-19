package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_convertToTitle(t *testing.T) {
	require.Equal(t, "A", convertToTitle(1))
	require.Equal(t, "AB", convertToTitle(28))
	require.Equal(t, "ZY", convertToTitle(701))
	require.Equal(t, "FXSHRXW", convertToTitle(2147483647))
}
