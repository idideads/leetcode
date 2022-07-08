package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isUgly(t *testing.T) {
	require.Equal(t, true, isUgly(1))
	require.Equal(t, true, isUgly(6))
	require.Equal(t, false, isUgly(14))
}
