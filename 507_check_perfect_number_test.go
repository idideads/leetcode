package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkPerfectNumber(t *testing.T) {
	require.Equal(t, false, checkPerfectNumber(1))
	require.Equal(t, true, checkPerfectNumber(28))
	require.Equal(t, false, checkPerfectNumber(7))
}
