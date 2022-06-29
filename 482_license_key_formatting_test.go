package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_licenseKeyFormatting(t *testing.T) {
	require.Equal(t, "5F3Z-2E9W", licenseKeyFormatting("5F3Z-2e-9-w", 4))
	require.Equal(t, "5F-3Z-2E-9W", licenseKeyFormatting("5F3Z-2e-9-w", 2))
	require.Equal(t, "2-5G-3J", licenseKeyFormatting("2-5g-3-J", 2))
	require.Equal(t, "AA-AA", licenseKeyFormatting("-a-a-a-a--", 2))
}

func Test_licenseKeyFormattingReverse(t *testing.T) {
	require.Equal(t, "5F3Z-2E9W", licenseKeyFormattingReverse("5F3Z-2e-9-w", 4))
	require.Equal(t, "5F-3Z-2E-9W", licenseKeyFormattingReverse("5F3Z-2e-9-w", 2))
	require.Equal(t, "2-5G-3J", licenseKeyFormattingReverse("2-5g-3-J", 2))
	require.Equal(t, "AA-AA", licenseKeyFormattingReverse("-a-a-a-a--", 2))
}
