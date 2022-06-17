package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_inorderTraversal(t *testing.T) {
	root := &TreeNode{1,
		&TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, &TreeNode{8, nil, nil}, &TreeNode{9, nil, nil}}},
		&TreeNode{3, &TreeNode{6, &TreeNode{10, nil, nil}, &TreeNode{11, nil, nil}}, &TreeNode{7, nil, nil}},
	}
	gotPath := inorderTraversal(root)
	wantPath := []int{4, 2, 8, 5, 9, 1, 10, 6, 11, 3, 7}
	require.NotNil(t, gotPath)
	require.EqualValues(t, wantPath, gotPath)
}
