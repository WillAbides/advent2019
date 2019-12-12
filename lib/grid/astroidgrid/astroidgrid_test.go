package astroidgrid

import (
	"fmt"
	"testing"

	"github.com/WillAbides/advent2019/lib/grid"
	"github.com/stretchr/testify/assert"
)

func TestGrid_isInGrid(t *testing.T) {
	g := Grid{
		northWest: northWest,
		southEast: grid.NewPoint(9, -9),
	}
	assert.True(t, g.containsPoint(grid.NewPoint(5, -5)))
	assert.True(t, g.containsPoint(grid.NewPoint(0, -5)))
	assert.True(t, g.containsPoint(grid.NewPoint(0, 0)))
	assert.True(t, g.containsPoint(grid.NewPoint(9, 0)))
	assert.True(t, g.containsPoint(grid.NewPoint(9, -9)))
	assert.False(t, g.containsPoint(grid.NewPoint(10, -9)))
	assert.False(t, g.containsPoint(grid.NewPoint(-1, -9)))
	assert.False(t, g.containsPoint(grid.NewPoint(5, -10)))
	assert.False(t, g.containsPoint(grid.NewPoint(5, 1)))
}

// Get all prime factors of a given number n
func PrimeFactors(n uint) []uint {
	factors := make([]uint, 0, n/2)
	if n < 2 {
		return factors
	}

	for i := uint(2); i*i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n = n / i
		}
	}

	if n > 1 {
		factors = append(factors, n)
	}

	return factors
}

func ReduceRatio(numerator, denominator int) (int, int) {
	negN := 1
	negD := 1
	if numerator < 0 {
		negN = -1
		numerator *= negN
	}
	if denominator < 0 {
		negD = -1
		denominator *= negD
	}
	if numerator == 0 {
		denominator = 1
	}
	if denominator == 0 {
		numerator = 1
	}
	nFactors := PrimeFactors(uint(numerator))
	dFactors := PrimeFactors(uint(denominator))
factorloop:
	for _, nFactor := range nFactors {
		for _, dFactor := range dFactors {
			if nFactor == dFactor {
				factor := int(nFactor)
				numerator, denominator = ReduceRatio(numerator/factor, denominator/factor)
				break factorloop
			}
		}
	}
	numerator *= negN
	denominator *= negD
	return numerator, denominator
}

func TestBlockingPositions(t *testing.T) {
	tests := []struct {
		southEast *grid.Point
		origin    *grid.Point
		blocker   *grid.Point
		want      []*grid.Point
	}{
		{
			southEast: &grid.Point{X: 9, Y: -9},
			origin:    &grid.Point{X: 0, Y: 0},
			blocker:   &grid.Point{X: 0, Y: 0},
			want:      []*grid.Point{},
		},
		{
			southEast: &grid.Point{X: 9, Y: -9},
			origin:    &grid.Point{X: 0, Y: 0},
			blocker:   &grid.Point{X: 3, Y: -3},
			want: []*grid.Point{
				{4, -4},
				{5, -5},
				{6, -6},
				{7, -7},
				{8, -8},
				{9, -9},
			},
		},
		{
			southEast: &grid.Point{X: 9, Y: -9},
			origin:    &grid.Point{X: 0, Y: 0},
			blocker:   &grid.Point{X: 3, Y: -2},
			want: []*grid.Point{
				{6, -4},
				{9, -6},
			},
		},
		{
			southEast: grid.NewPoint(9, -9),
			origin:    grid.NewPoint(0, 0),
			blocker:   grid.NewPoint(3, 0),
			want: []*grid.Point{
				grid.NewPoint(4, 0),
				grid.NewPoint(5, 0),
				grid.NewPoint(6, 0),
				grid.NewPoint(7, 0),
				grid.NewPoint(8, 0),
				grid.NewPoint(9, 0),
			},
		},
		{
			southEast: grid.NewPoint(9, -9),
			origin:    grid.NewPoint(9, -9),
			blocker:   grid.NewPoint(3, -9),
			want: []*grid.Point{
				grid.NewPoint(2, -9),
				grid.NewPoint(1, -9),
				grid.NewPoint(0, -9),
			},
		},
	}

	for _, td := range tests {
		t.Run("", func(t *testing.T) {
			got := BlockingPositions(td.southEast, td.origin, td.blocker)
			assert.ElementsMatch(t, td.want, got)
		})
	}
}

var exGrid = `
.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##
`

func TestLoadRoidGrid(t *testing.T) {
	grd := LoadRoidGrid(exGrid)
	fmt.Println(len(grd.roidGrid))
}

func TestGrid_markPositionBlocked(t *testing.T) {
	grd := LoadRoidGrid(exGrid)
	roid := grd.roidGrid[*grid.NewPoint(1,0)].(*Astroid)
	fmt.Println(*roid)
	grd.markPositionBlocked(grid.NewPoint(1,0))
	roid = grd.roidGrid[*grid.NewPoint(1,0)].(*Astroid)
	fmt.Println(*roid)
	grd.Reset()
	roid = grd.roidGrid[*grid.NewPoint(1,0)].(*Astroid)
	fmt.Println(*roid)
}

func TestGrid_markBlockedRoids(t *testing.T) {
	grd := LoadRoidGrid(exGrid)
	grd.markBlockedRoids(grid.NewPoint(1,0), grid.NewPoint(4,0))
	roid := grd.roidGrid[*grid.NewPoint(4,0)].(*Astroid)
	fmt.Println(*roid)
	roid = grd.roidGrid[*grid.NewPoint(5,0)].(*Astroid)
	fmt.Println(*roid)
	roid = grd.roidGrid[*grid.NewPoint(19,0)].(*Astroid)
	fmt.Println(*roid)
}

func TestFoo(t *testing.T) {
	grd := LoadRoidGrid(exGrid)
	grd.markBlocked(grid.NewPoint(11, -13))
	fmt.Println(grd.countVisible() - 1)
}

func TestFindBestSpot(t *testing.T) {
	grd := LoadRoidGrid(exGrid)
	wantCount := 210
	wantPos := grid.NewPoint(11, -13)
	gotPoint, gotCount := grd.FindBestSpot()
	assert.Equal(t, *wantPos, gotPoint)
	assert.Equal(t, wantCount, gotCount)
}
