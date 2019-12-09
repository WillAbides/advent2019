package day8

import (
	"strings"
	"testing"

	"github.com/WillAbides/advent2019/lib"
	"github.com/WillAbides/advent2019/lib/imageviewer"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	digits := lib.StringDigits(string(lib.MustReadFile("input.txt")))
	image := imageviewer.New(6, 25, digits)
	assert.Equal(t, 1452, image.Checksum())
}

func TestPart2(t *testing.T) {
	want := `
◻️◻️◻️◼️◼️◻️◼️◼️◻️◼️◻️◻️◻️◼️◼️◻️◻️◻️◻️◼️◻️◼️◼️◻️◼️
◻️◼️◼️◻️◼️◻️◼️◼️◻️◼️◻️◼️◼️◻️◼️◻️◼️◼️◼️◼️◻️◼️◼️◻️◼️
◻️◼️◼️◻️◼️◻️◻️◻️◻️◼️◻️◼️◼️◻️◼️◻️◻️◻️◼️◼️◻️◼️◼️◻️◼️
◻️◻️◻️◼️◼️◻️◼️◼️◻️◼️◻️◻️◻️◼️◼️◻️◼️◼️◼️◼️◻️◼️◼️◻️◼️
◻️◼️◼️◼️◼️◻️◼️◼️◻️◼️◻️◼️◼️◼️◼️◻️◼️◼️◼️◼️◻️◼️◼️◻️◼️
◻️◼️◼️◼️◼️◻️◼️◼️◻️◼️◻️◼️◼️◼️◼️◻️◻️◻️◻️◼️◼️◻️◻️◼️◼️`
	want = strings.TrimSpace(want)
	digits := lib.StringDigits(string(lib.MustReadFile("input.txt")))
	image := imageviewer.New(6, 25, digits)
	assert.Equal(t, want, image.Render())
}
