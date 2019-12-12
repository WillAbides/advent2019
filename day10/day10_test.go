package day10

import (
	"testing"

	"github.com/WillAbides/advent2019/lib"
	"github.com/WillAbides/advent2019/lib/grid/astroidgrid"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	grd := astroidgrid.LoadRoidGrid(string(lib.MustReadFile("input.txt")))
	_, count := grd.FindBestSpot()
	assert.Equal(t, 260, count)
}
