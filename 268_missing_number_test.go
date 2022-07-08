package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_missingNumber(t *testing.T) {
	require.Equal(t, 2, missingNumber([]int{3, 0, 1}))
}
