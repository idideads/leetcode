package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverseString(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	require.EqualValues(t, []byte{'o', 'l', 'l', 'e', 'h'}, s)
}
