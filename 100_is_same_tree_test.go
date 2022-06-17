package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isSameTree(t *testing.T) {
	root1 := &TreeNode{1,
		&TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, &TreeNode{8, nil, nil}, &TreeNode{9, nil, nil}}},
		&TreeNode{3, &TreeNode{6, &TreeNode{10, nil, nil}, &TreeNode{11, nil, nil}}, &TreeNode{7, nil, nil}},
	}
	root2 := &TreeNode{1,
		&TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, &TreeNode{8, nil, nil}, &TreeNode{9, nil, nil}}},
		&TreeNode{3, &TreeNode{6, &TreeNode{10, nil, nil}, &TreeNode{11, nil, nil}}, &TreeNode{7, nil, nil}},
	}
	require.Equal(t, true,isSameTree(root1, root2))
}
