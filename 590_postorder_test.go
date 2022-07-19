package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_postorder(t *testing.T) {
	n6 := &Node{Val: 6}
	n5 := &Node{Val: 5}
	n3 := &Node{Val: 3, Children: []*Node{n5, n6}}
	n2 := &Node{Val: 2}
	n4 := &Node{Val: 4}
	n1 := &Node{Val: 1, Children: []*Node{n3, n2, n4}}
	require.EqualValues(t, []int{5, 6, 3, 2, 4, 1}, postorder(n1))
}
