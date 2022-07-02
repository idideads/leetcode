package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findWords(t *testing.T) {
	require.EqualValues(t, []string{"Alaska", "Dad"}, findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
	require.EqualValues(t, []string{}, findWords([]string{"omk"}))
	require.EqualValues(t, []string{"adsdf", "sfd"}, findWords([]string{"adsdf", "sfd"}))
	require.EqualValues(t, []string{"Aasdfghjkl", "Qwertyuiop", "zZxcvbnm"}, findWords([]string{"Aasdfghjkl", "Qwertyuiop", "zZxcvbnm"}))
}
