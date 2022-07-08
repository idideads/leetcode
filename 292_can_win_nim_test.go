package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canWinNim(t *testing.T) {
	require.Equal(t, true, canWinNim(25))
}
