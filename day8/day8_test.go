package day8

import (
	"fmt"
	"strings"
	"testing"

	"advent2019/lib"
)

func rune2int(r rune) int {
	i := int(r - 48)
	if i < 0 || i >  9 {
		panic("unexpected rune")
	}
	return i
}

func digitList(input string) []int {
	result := make([]int, len(input))
	for i, r := range input {
		result[i] = rune2int(r)
	}
	return result
}

func buildLayers(height, width int, digits []int) [][]int {
	layerSize := height * width
	result := make([][]int, 0)
	for i := 0; i < len(digits) - 1; i += layerSize {
		result = append(result,  digits[i:i + layerSize])
	}
	return result
}

func countDigits(layer []int, digit int) int {
	count := 0
	for _, i := range layer {
		if i == digit {
			count++
		}
	}
	return count
}

func layerWithFewestZeros(layers [][]int) []int {
	minZeros := -1
	targetLayer := -1
	for i, layer := range layers {
		cnt := countDigits(layer, 0)
		if minZeros == -1 {
			minZeros = cnt
			targetLayer = i
		}
		if minZeros > cnt {
			minZeros = cnt
			targetLayer = i
		}
	}
	return layers[targetLayer]
}

func visibleLayer(layers [][]int) []int {
	layerLen := len(layers[0])
	result := make([]int, layerLen)
	for i := 0; i < layerLen; i++ {
		for _, layer := range layers {
			pixel := layer[i]
			if pixel < 2 {
				result[i] = pixel
				break
			}
		}
	}
	return result
}

const width = 25
const height = 6

func TestPart2(t *testing.T) {
	input := string(lib.MustReadFile("input.txt"))
	input = strings.TrimSpace(input)
	digits := digitList(input)
	layers := buildLayers(height, width, digits)
	vis := visibleLayer(layers)
	fmt.Println(vis)
	cursor := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			switch vis[cursor] {
			case 0:
				fmt.Print("◼️")
			case 1:
				fmt.Print("◻️")
			}
			cursor++
		}
		fmt.Println("")
	}
}

func TestPart1(t *testing.T) {
	input := string(lib.MustReadFile("input.txt"))
	input = strings.TrimSpace(input)
	digits := digitList(input)
	l := buildLayers(height, width, digits)
	targetLayer := layerWithFewestZeros(l)
	ones := countDigits(targetLayer, 1)
	twos := countDigits(targetLayer, 2)
	fmt.Println(ones * twos)
}

