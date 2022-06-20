package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_containsDuplicate(t *testing.T) {
	require.Equal(t, true, containsDuplicate([]int{1, 2, 3, 5, 7, 11, 13, 17, 5}))
	require.Equal(t, false, containsDuplicate([]int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29}))
}
