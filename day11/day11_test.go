package day11

import (
	"strings"
	"testing"

	"github.com/WillAbides/advent2019/lib"
	"github.com/WillAbides/advent2019/lib/grid"
	"github.com/WillAbides/advent2019/lib/paintbot"
	"github.com/WillAbides/advent2019/lib/paintgrid"
	"github.com/stretchr/testify/assert"
)

var realProgram = lib.CSInts(string(lib.MustReadFile("input.txt")))

func TestPart1(t *testing.T) {
	bot := paintbot.New(realProgram)
	panel := paintgrid.New()
	panel.SetDefaultColor(paintgrid.Black)
	bot.PlaceOnPanel(panel, grid.Point{X: 0, Y: 0}, grid.North)
	err := bot.Run()
	assert.NoError(t, err)
	paintedPanels := panel.PaintedCount()
	assert.Equal(t, 2054, paintedPanels)
}

func TestPart2(t *testing.T) {
	bot := paintbot.New(realProgram)
	panel := paintgrid.New()
	panel.SetDefaultColor(paintgrid.Black)
	panel.Paint(grid.Point{X: 0, Y: 0}, paintgrid.White)
	bot.PlaceOnPanel(panel, grid.Point{X: 0, Y: 0}, grid.North)
	err := bot.Run()
	assert.NoError(t, err)
	image := panel.Image()
		want := `
◻️◼️◻️◻️◼️◻️◼️◼️◼️◻️◻️◼️◼️◼️◼️◻️◼️◼️◼️◼️◻️◻️◼️◼️◻️◻️◻️◻️◼️◼️◻️◼️◻️◻️◼️◻️◼️◼️◼️◻️◻️◻️◻️
◻️◼️◻️◼️◻️◻️◼️◻️◻️◼️◻️◻️◻️◻️◼️◻️◼️◻️◻️◻️◻️◼️◻️◻️◼️◻️◻️◻️◻️◼️◻️◼️◻️◻️◼️◻️◼️◻️◻️◼️◻️◻️◻️
◻️◼️◼️◻️◻️◻️◼️◻️◻️◼️◻️◻️◻️◼️◻️◻️◼️◼️◼️◻️◻️◼️◻️◻️◼️◻️◻️◻️◻️◼️◻️◼️◼️◼️◼️◻️◼️◼️◼️◻️◻️◻️◻️
◻️◼️◻️◼️◻️◻️◼️◼️◼️◻️◻️◻️◼️◻️◻️◻️◼️◻️◻️◻️◻️◼️◼️◼️◼️◻️◻️◻️◻️◼️◻️◼️◻️◻️◼️◻️◼️◻️◻️◼️◻️◻️◻️
◻️◼️◻️◼️◻️◻️◼️◻️◼️◻️◻️◼️◻️◻️◻️◻️◼️◻️◻️◻️◻️◼️◻️◻️◼️◻️◼️◻️◻️◼️◻️◼️◻️◻️◼️◻️◼️◻️◻️◼️◻️◻️◻️
◻️◼️◻️◻️◼️◻️◼️◻️◻️◼️◻️◼️◼️◼️◼️◻️◼️◼️◼️◼️◻️◼️◻️◻️◼️◻️◻️◼️◼️◻️◻️◼️◻️◻️◼️◻️◼️◼️◼️◻️◻️◻️◻️`

		want = strings.TrimSpace(want)
		assert.Equal(t, want, image.NegativeImage().Render())
}
