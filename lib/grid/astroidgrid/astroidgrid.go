package astroidgrid

import (
	"strings"

	"github.com/WillAbides/advent2019/lib"
	"github.com/WillAbides/advent2019/lib/grid"
)

var northWest = &grid.Point{
	X: 0, Y: 0,
}

type Grid struct {
	northWest *grid.Point
	southEast *grid.Point
	roidGrid  grid.ValueGrid
}

func (g *Grid) containsPoint(pos *grid.Point) bool {
	return pos.X <= g.maxX() &&
		pos.X >= g.minX() &&
		pos.Y <= g.maxY() &&
		pos.Y >= g.minY()

}

func (g *Grid) maxX() int {
	return g.southEast.X
}

func (g *Grid) minX() int {
	return g.northWest.X
}

func (g *Grid) maxY() int {
	return g.northWest.Y
}

func (g *Grid) minY() int {
	return g.southEast.Y
}

func BlockingPositions(southEast *grid.Point, origin *grid.Point, blocker *grid.Point) []*grid.Point {
	grd := &Grid{
		northWest: northWest,
		southEast: southEast,
	}
	var result []*grid.Point

	return recursiveBlockingPositions(result, grd, origin, blocker)
}

func (g *Grid) blockingPositions(origin *grid.Point, blocker *grid.Point) []*grid.Point {
	var result []*grid.Point
	return recursiveBlockingPositions(result, g, origin, blocker)
}

func recursiveBlockingPositions(blocked []*grid.Point, grd *Grid, origin, blocker *grid.Point) []*grid.Point {
	if *origin == *blocker {
		return blocked
	}
	if ! grd.containsPoint(origin) || ! grd.containsPoint(blocker) {
		return blocked
	}
	xDifferential := origin.X - blocker.X
	yDifferential := origin.Y - blocker.Y
	xDifferential, yDifferential = lib.ReduceRatio(xDifferential, yDifferential)
	blocker = &grid.Point{
		X: blocker.X - xDifferential,
		Y: blocker.Y - yDifferential,
	}
	if grd.containsPoint(blocker) {
		blocked = append(blocked, blocker)
		return recursiveBlockingPositions(blocked, grd, origin, blocker)
	}
	return blocked
}

func LoadRoidGrid(input string) *Grid {
	grd := &Grid{
		northWest: northWest,
		roidGrid:  grid.ValueGrid{},
	}
	input = strings.TrimSpace(input)
	for i, line := range strings.Split(input, "\n") {
		y := 0 - i
		for x, ch := range line {
			pos := grid.NewPoint(x, y)
			grd.southEast = pos
			if ch == '#' {
				grd.roidGrid[*pos] = &Astroid{}
			}
		}
	}
	return grd
}

func (g *Grid) Reset() {
	for _, val := range g.roidGrid {
		if val == nil {
			continue
		}
		roid, ok := val.(*Astroid)
		if !ok {
			continue
		}
		roid.blocked = false
	}
}

func (g *Grid) markBlockedRoids(origin, blocker *grid.Point) {
	for _, point := range g.blockingPositions(origin, blocker) {
		g.markPositionBlocked(point)
	}
}

func (g *Grid) markBlocked(origin *grid.Point) {
	for point, v := range g.roidGrid {
		roid, ok := v.(*Astroid)
		if !ok {
			continue
		}
		if roid.blocked {
			continue
		}
		if point == *origin {
			continue
		}
		g.markBlockedRoids(origin, &point)
	}
}

func (g *Grid) markPositionBlocked(pos *grid.Point) {
	val := g.roidGrid[*pos]
	if val == nil {
		return
	}
	roid, ok := val.(*Astroid)
	if !ok {
		return
	}
	roid.blocked = true
}

func (g *Grid) countVisible() int {
	count := 0
	for _, v := range g.roidGrid {
		roid, ok := v.(*Astroid)
		if !ok {
			continue
		}
		if ! roid.blocked {
			count++
		}
	}
	return count
}

func (g *Grid) FindBestSpot() (grid.Point, int) {
	maxVisible := 0
	var bestPos grid.Point
	for point := range g.roidGrid {
		g.markBlocked(&point)
		vis := g.countVisible() - 1
		if maxVisible < vis {
			maxVisible = vis
			bestPos = point
		}
		g.Reset()
	}
	return bestPos, maxVisible
}

type Astroid struct {
	blocked bool
}
