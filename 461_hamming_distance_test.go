package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_hammingDistance(t *testing.T) {
	require.Equal(t, 2, hammingDistance(1, 4))
	require.Equal(t, 1, hammingDistance(1, 3))
	require.Equal(t, 3, hammingDistance(6, 8))
}
