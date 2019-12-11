package paintgrid

import (
	"github.com/WillAbides/advent2019/lib/grid"
	"github.com/WillAbides/advent2019/lib/imageviewer"
)

type Color int

const defaultDefaultColor = NoColor

const (
	Black   Color = 0
	White   Color = 1
	NoColor Color = 2
)

type Grid struct {
	valueGrid     grid.ValueGrid
	_defaultColor *Color
}

func New() *Grid {
	return &Grid{
		valueGrid: grid.ValueGrid{},
	}
}

func (p *Grid) defaultColor() Color {
	if p._defaultColor != nil {
		return *p._defaultColor
	}
	return defaultDefaultColor
}

func (p *Grid) SetDefaultColor(defaultColor Color) {
	p._defaultColor = &defaultColor
}

func (p *Grid) Paint(position grid.Point, color Color) {
	p.valueGrid[position] = color
}

func (p *Grid) GetColor(position grid.Point) Color {
	val := p.valueGrid[position]
	if val == nil {
		val = p.defaultColor()
	}
	return val.(Color)
}

func (p *Grid) Image() *imageviewer.Image {
	rows := p.valueGrid.Rows()
	height := p.valueGrid.Height()
	width := p.valueGrid.Width()
	imageData := make([]int, 0, height * width)
	for _, row := range rows {
		for _, val := range row {
			if val == nil {
				val = p.defaultColor()
			}
			imageData = append(imageData, int(val.(Color)))
		}
	}
	return imageviewer.New(height, width, imageData)
}

func (p *Grid) PaintedCount() int {
	return len(p.valueGrid)
}
