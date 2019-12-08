package imageviewer

import (
	"strings"
	"testing"

	"advent2019/lib"

	"github.com/stretchr/testify/assert"
)

func TestImage_NegativeImage(t *testing.T) {
	digits := lib.StringDigits(string(lib.MustReadFile("input.txt")))
	image := New(6, 25, digits)
	got := image.NegativeImage().Render()
	want := `
◼️◼️◼️◻️◻️◼️◻️◻️◼️◻️◼️◼️◼️◻️◻️◼️◼️◼️◼️◻️◼️◻️◻️◼️◻️
◼️◻️◻️◼️◻️◼️◻️◻️◼️◻️◼️◻️◻️◼️◻️◼️◻️◻️◻️◻️◼️◻️◻️◼️◻️
◼️◻️◻️◼️◻️◼️◼️◼️◼️◻️◼️◻️◻️◼️◻️◼️◼️◼️◻️◻️◼️◻️◻️◼️◻️
◼️◼️◼️◻️◻️◼️◻️◻️◼️◻️◼️◼️◼️◻️◻️◼️◻️◻️◻️◻️◼️◻️◻️◼️◻️
◼️◻️◻️◻️◻️◼️◻️◻️◼️◻️◼️◻️◻️◻️◻️◼️◻️◻️◻️◻️◼️◻️◻️◼️◻️
◼️◻️◻️◻️◻️◼️◻️◻️◼️◻️◼️◻️◻️◻️◻️◼️◼️◼️◼️◻️◻️◼️◼️◻️◻️`
	want = strings.TrimSpace(want)
	assert.Equal(t, want , strings.TrimSpace(got))
}

func TestImage_Render(t *testing.T) {
	digits := lib.StringDigits(string(lib.MustReadFile("input.txt")))
	image := New(6, 25, digits)
	got := image.Render()
	want := `
◻️◻️◻️◼️◼️◻️◼️◼️◻️◼️◻️◻️◻️◼️◼️◻️◻️◻️◻️◼️◻️◼️◼️◻️◼️
◻️◼️◼️◻️◼️◻️◼️◼️◻️◼️◻️◼️◼️◻️◼️◻️◼️◼️◼️◼️◻️◼️◼️◻️◼️
◻️◼️◼️◻️◼️◻️◻️◻️◻️◼️◻️◼️◼️◻️◼️◻️◻️◻️◼️◼️◻️◼️◼️◻️◼️
◻️◻️◻️◼️◼️◻️◼️◼️◻️◼️◻️◻️◻️◼️◼️◻️◼️◼️◼️◼️◻️◼️◼️◻️◼️
◻️◼️◼️◼️◼️◻️◼️◼️◻️◼️◻️◼️◼️◼️◼️◻️◼️◼️◼️◼️◻️◼️◼️◻️◼️
◻️◼️◼️◼️◼️◻️◼️◼️◻️◼️◻️◼️◼️◼️◼️◻️◻️◻️◻️◼️◼️◻️◻️◼️◼️`
	want = strings.TrimSpace(want)
	assert.Equal(t, want , strings.TrimSpace(got))
}

func TestImage_renderPixels(t *testing.T) {
	data := []int{0, 2, 2, 2, 1, 1, 2, 2, 2, 2, 1, 2, 0, 0, 0, 0}
	want := []pixel{0, 1, 1, 0}
	image := New(2, 2, data)
	got := image.renderPixels()
	assert.Equal(t, want, got)
}

func TestImage_pixelCounts(t *testing.T) {
	image := New(2, 2, []int{0, 0, 1, 2})
	want := map[pixel]int{
		0: 2,
		1: 1,
		2: 1,
	}
	got := image.pixelCounts()
	assert.Equal(t, want, got)
}

func TestImage_Checksum(t *testing.T) {
	digits := lib.StringDigits(string(lib.MustReadFile("input.txt")))
	image := New(6, 25, digits)
	got := image.Checksum()
	want := 1452
	assert.Equal(t, want, got)
}
