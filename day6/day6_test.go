package day6

import (
	"strings"
	"testing"

	"github.com/WillAbides/advent2019/lib"
	"github.com/WillAbides/advent2019/lib/orbitcalc"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := string(lib.MustReadFile("input.txt"))
	input = strings.TrimSpace(input)
	got := orbitcalc.OrbitCount(input)
	assert.Equal(t, 223251, got)
}

func TestPart2(t *testing.T) {
	input := string(lib.MustReadFile("input.txt"))
	got := orbitcalc.CalcOrbitTransfers(input, "YOU", "SAN")
	assert.Equal(t, 430, got)
}
