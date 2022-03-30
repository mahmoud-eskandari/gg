package gg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	is := assert.New(t)
	is.Equal(Unique([]int{8, 7, 6, 5, 4, 0, 0}), []int{8, 7, 6, 5, 4, 0})
	is.Equal(Unique([]int{8, 8, 7, 6, 5, 4, 0, 0}), []int{8, 7, 6, 5, 4, 0})

}
