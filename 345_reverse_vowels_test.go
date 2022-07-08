package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverseVowels(t *testing.T) {
	require.Equal(t, "leotcede", reverseVowels("leetcode"))
}
