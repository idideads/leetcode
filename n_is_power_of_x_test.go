package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNisPowerOfX(t *testing.T) {
	require.Equal(t, true, NisPowerOfX(16, 4))
	require.Equal(t, true, NisPowerOfX(8, 2))
	require.Equal(t, false, NisPowerOfX(20, 3))
}
