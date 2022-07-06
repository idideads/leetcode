package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_arrayPairSum(t *testing.T) {
	require.Equal(t, 4, arrayPairSum([]int{1, 4, 3, 2}))
	require.Equal(t, 9, arrayPairSum([]int{6, 2, 6, 5, 1, 2}))
}
