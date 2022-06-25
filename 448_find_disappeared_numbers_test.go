package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findDisappearedNumbers(t *testing.T) {
	require.NotNil(t, findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	require.EqualValues(t, []int{5, 6}, findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
