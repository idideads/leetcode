package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countBits(t *testing.T) {
	require.EqualValues(t, []int{0, 1, 1}, countBits(2))
	require.EqualValues(t, []int{0, 1, 1, 2, 1, 2}, countBits(5))
}
