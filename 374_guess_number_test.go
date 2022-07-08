package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_guessNumber(t *testing.T) {
	require.Equal(t, 4, guessNumber(4))
	require.Equal(t, 6, guessNumber(6))
	require.Equal(t, 6, guessNumber(13))
}
