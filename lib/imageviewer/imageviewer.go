package imageviewer

import (
	"bytes"
)

type pixel int

const (
	pixelTransparent pixel = 2
	pixelWhite       pixel = 1
	pixelBlack       pixel = 0
)

var pixelStrings = map[pixel]string {
	pixelWhite: "◻️",
	pixelBlack: "◼️",
}

type Image struct {
	height, width int
	data          []int
}

func New(height, width int, data []int) *Image {
	return &Image{
		height: height,
		width:  width,
		data:   data,
	}
}

//Checksum is :
//find the layer that contains the fewest 0 digits. On that layer, what is the number of 1 digits multiplied by the
//number of 2 digits?
func (m *Image) Checksum() int {
	layers := m.layers()
	minZeros := -1
	targetLayer := -1
	for i, layer := range layers {
		zeroCount := layer.pixelCounts()[0]
		if minZeros == -1 || minZeros > zeroCount {
			minZeros = zeroCount
			targetLayer = i
		}
	}
	counts := layers[targetLayer].pixelCounts()
	return counts[1] * counts[2]
}

//NegativeImage swaps black and white pixels
func (m *Image) NegativeImage() *Image {
	image := &Image{
		height: m.height,
		width: m.width,
		data: make([]int, len(m.data)),
	}
	copy(image.data, m.data)
	for i, p := range image.data {
		switch p {
		case int(pixelWhite):
			image.data[i] = int(pixelBlack)
		case int(pixelBlack):
			image.data[i] = int(pixelWhite)
		}
	}
	return image
}

//Render returns a string with the rendered image
func (m *Image) Render() string {
	var buf bytes.Buffer
	pixels := m.renderPixels()
	for i, p := range pixels {
		if i != 0 && i % m.width == 0 {
			_, _ = buf.WriteRune('\n')
		}
		_, _ = buf.WriteString(pixelStrings[p])
	}
	return buf.String()
}

func (m *Image) pixels() []pixel {
	result := make([]pixel, len(m.data))
	for i, d := range m.data {
		result[i] = pixel(d)
	}
	return result
}

func (m *Image) layerSize() int {
	return m.height * m.width
}

func (m *Image) layerCount() int {
	return len(m.data) / m.layerSize()
}

func (m *Image) layers() []*Image {
	result := make([]*Image, m.layerCount())
	data := make([]int, len(m.data))
	copy(data, m.data)
	for i := 0; i < m.layerCount(); i++ {
		result[i] = New(m.height, m.width, data[:m.layerSize()])
		data = data[m.layerSize():]
	}
	return result
}

func (m *Image) pixelCounts() map[pixel]int {
	result := map[pixel]int{}
	for _, p := range m.renderPixels() {
		result[p]++
	}
	return result
}

func (m *Image) renderPixels() []pixel {
	if m.layerCount() < 2 {
		return m.pixels()
	}
	layers := make([][]pixel, m.layerCount())
	for i, l := range m.layers() {
		layers[i] = l.renderPixels()
	}
	pixels := make([]pixel, m.layerSize())
	for i := range pixels {
		for _, lp := range layers {
			if lp[i] != pixelTransparent {
				pixels[i] = lp[i]
				break
			}
		}
	}
	return pixels
}
