package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isHappy(t *testing.T) {
	require.Equal(t, true, isHappy(19))
	require.Equal(t, false, isHappy(2))
}
