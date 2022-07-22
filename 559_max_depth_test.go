package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxDepthTree(t *testing.T) {
	n9 := &Node{Val: 9}
	n8 := &Node{Val: 8, Children: []*Node{n9}}
	n7 := &Node{Val: 7}
	n6 := &Node{Val: 6}
	n5 := &Node{Val: 5}
	n3 := &Node{Val: 3, Children: []*Node{n5, n6}}
	n2 := &Node{Val: 2}
	n4 := &Node{Val: 4, Children: []*Node{n7, n8}}
	n1 := &Node{Val: 1, Children: []*Node{n3, n2, n4}}

	require.Equal(t, 1, maxDepthTree(n2))
	require.Equal(t, 1, maxDepthTree(n6))
	require.Equal(t, 4, maxDepthTree(n1))
}
