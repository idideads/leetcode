package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isValid(t *testing.T) {
	require.Equal(t, true, isValid("()[]{}"))
}
