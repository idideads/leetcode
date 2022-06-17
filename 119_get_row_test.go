package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getRow(t *testing.T) {
	got := getRow(4)
	require.NotNil(t, got)
	require.EqualValues(t, []int{1, 4, 6, 4, 1}, got)
}
