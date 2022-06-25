package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_arrangeCoins(t *testing.T) {
	require.Equal(t, 1, arrangeCoins(1))
	require.Equal(t, 2, arrangeCoins(5))
	require.Equal(t, 3, arrangeCoins(8))
	require.Equal(t, 5, arrangeCoins(15))
}
