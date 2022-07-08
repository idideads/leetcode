package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_firstBadVersion(t *testing.T) {
	require.Equal(t, 6, firstBadVersion(8))
	require.Equal(t, 6, firstBadVersion(22))
}
