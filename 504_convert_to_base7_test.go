package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_convertToBase7(t *testing.T) {
	require.Equal(t, "0", convertToBase7(0))
	require.Equal(t, "1", convertToBase7(1))
	require.Equal(t, "6", convertToBase7(6))
	require.Equal(t, "10", convertToBase7(7))
	require.Equal(t, "-10", convertToBase7(-7))
	require.Equal(t, "100", convertToBase7(49))
	require.Equal(t, "202", convertToBase7(100))
}
