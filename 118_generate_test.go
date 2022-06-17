package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_generate(t *testing.T) {
	got := generate(5)
	require.NotNil(t, got)
	want := [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}
	require.EqualValues(t, want, got)
}
