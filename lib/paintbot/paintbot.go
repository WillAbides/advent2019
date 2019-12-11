package paintbot

import (
	"github.com/WillAbides/advent2019/lib/grid"
	"github.com/WillAbides/advent2019/lib/intcomputer"
	"github.com/WillAbides/advent2019/lib/paintgrid"
)

type Bot struct {
	program     []int64
	computer    *intcomputer.IntComputer
	orientation grid.Orientation
	point       grid.Point
	outputCount int64
	grid        *paintgrid.Grid
}

func (p *Bot) turn(direction grid.RelativeDirection) {
	p.orientation = grid.Turn(p.orientation, direction)
}

func (p *Bot) move() {
	p.point = grid.Move(p.point, p.orientation, 1)
}

func (p *Bot) currentPanel() paintgrid.Color {
	return p.grid.GetColor(p.point)
}

func (p *Bot) paint(color paintgrid.Color) {
	p.grid.Paint(p.point, color)
}

func (p *Bot) computerOutputHandler(c *intcomputer.IntComputer, n int64) error {
	outputType := p.outputCount % 2
	switch outputType {
	case 0:
		var color paintgrid.Color
		switch n {
		case 0:
			color = paintgrid.Black
		case 1:
			color = paintgrid.White
		}
		p.paint(color)
	case 1:
		var dir grid.RelativeDirection
		switch n {
		case 0:
			dir = grid.Left
		case 1:
			dir = grid.Right
		}
		p.turn(dir)
		p.move()
	}
	p.outputCount++
	return nil
}

func (p *Bot) computerInputter() (int64, error) {
	return int64(p.currentPanel()), nil
}

func New(program []int64) *Bot {
	return &Bot{
		program: program,
	}
}

func (b *Bot) PlaceOnPanel(paintGrid *paintgrid.Grid, position grid.Point, orientation grid.Orientation) {
	b.grid = paintGrid
	b.point = position
	b.orientation = orientation
}

func (p *Bot) Run() error {
	p.computer = intcomputer.NewIntComputer(p.program, p.computerOutputHandler, p.computerInputter)
	return p.computer.RunOperations()
}
