package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findRestaurant(t *testing.T) {
	require.EqualValues(t, []string{"Shogun"}, findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}))
	require.EqualValues(t, []string{"Shogun"}, findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"KFC", "Shogun", "Burger King"}))
	require.EqualValues(t, []string{"bbb", "aaa"}, findRestaurant([]string{"aaa", "bbb", "ccc", "ddd"}, []string{"bbb", "aaa", "tttt"}))

}
