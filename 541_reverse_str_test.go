package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverseStr(t *testing.T) {
	require.Equal(t, "bacdfeg", reverseStr("abcdefg", 2))
	require.Equal(t, "bacd", reverseStr("abcd", 2))
	require.Equal(t, "dcba", reverseStr("abcd", 5))
	require.Equal(t, "fdcqkmxwholhytmhafpesaentdvxginrjlyqzyhehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqllgsqddebemjanqcqnfkjmi",
		reverseStr("hyzqyljrnigxvdtneasepfahmtyhlohwxmkqcdfehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqlimjkfnqcqnajmebeddqsgl", 39))

}
