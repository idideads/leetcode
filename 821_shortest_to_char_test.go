package leetcode

import (

	// "log"
	// "reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_shortestToChar(t *testing.T) {
	require.EqualValues(t, []int{0, 1, 1, 0}, shortestToChar("baab", 'b'))
	require.EqualValues(t, []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0, 1, 2, 3}, shortestToChar("loveleetcodeabc", 'e'))
}
