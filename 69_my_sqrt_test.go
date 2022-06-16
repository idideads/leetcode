package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mySqrt(t *testing.T) {
	require.Equal(t, 2, mySqrt(4))
	require.Equal(t, 2, mySqrt(8))
}
