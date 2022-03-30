package gg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexOf(t *testing.T) {
	is := assert.New(t)

	is.Equal(IndexOf([]int{8, 7, 6, 5, 4, 0}, 3), -1)
	is.Equal(IndexOf([]int{8, 7, 6, 5, 4, 0, 2, 2}, 2), 6)
	is.Equal(IndexOf([]int{}, 0), -1)
	is.Equal(IndexOf([]string{}, ""), -1)
	is.Equal(IndexOf([]string{"a"}, "A"), -1)
	is.Equal(IndexOf([]string{"a"}, "a"), 0)

}

func TestLastIndexOf(t *testing.T) {
	is := assert.New(t)

	is.Equal(LastIndexOf([]int{0, 1, 2, 1, 2, 3}, 2), 4)
	is.Equal(LastIndexOf([]string{"Hello", "b", "a", "a"}, "a"), 3)
	is.Equal(LastIndexOf([]string{}, ""), -1)
}

func TestAllIndexesOf(t *testing.T) {
	is := assert.New(t)

	is.Equal(AllIndexesOf([]int{0, 1, 2, 1, 2, 3}, 2), []int{2, 4})
	is.Equal(AllIndexesOf([]string{"Hello", "b", "a", "a"}, "a"), []int{2, 3})
	is.Equal(AllIndexesOf([]string{}, ""), []int{})
}

func TestCountByValue(t *testing.T) {
	is := assert.New(t)

	is.Equal(CountByValue([]int{0, 1, 2, 1, 2, 3}, 2), 2)
	is.Equal(CountByValue([]int{}, 2), 0)
	is.Equal(CountByValue([]string{"A", "b", "a"}, "A"), 1)
}
