package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findComplement(t *testing.T) {
	require.Equal(t, 2, findComplement(5))
	require.Equal(t, 0, findComplement(1))
	require.Equal(t, 0, findComplement(7))
	require.Equal(t, 1, findComplement(2))
}
