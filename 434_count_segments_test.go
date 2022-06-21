package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countSegments(t *testing.T) {
	require.Equal(t, 0, countSegments(""))
	require.Equal(t, 0, countSegments(" "))
	require.Equal(t, 5, countSegments("Hello, my name is John"))
	require.Equal(t, 3, countSegments("   aa  bb		cc  "))
}
