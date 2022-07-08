package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPerfectSquare(t *testing.T) {
	require.Equal(t, true, isPerfectSquare(4))
	require.Equal(t, false, isPerfectSquare(8))
	require.Equal(t, true, isPerfectSquare(19*19))
}
