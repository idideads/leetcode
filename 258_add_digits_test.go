package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_addDigits(t *testing.T) {
	require.Equal(t, 0, addDigits(0))
	require.Equal(t, 2, addDigits(38))
}
