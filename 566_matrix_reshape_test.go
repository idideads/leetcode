package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_matrixReshape(t *testing.T) {
	require.EqualValues(t, [][]int{{1, 2, 3, 4}}, matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4))
	require.EqualValues(t, [][]int{{1, 2}, {3, 4}}, matrixReshape([][]int{{1, 2}, {3, 4}}, 2, 4))
}
