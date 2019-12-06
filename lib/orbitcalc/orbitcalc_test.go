package orbitcalc

import (
	"strings"
	"testing"

	"advent2019/lib"

	"github.com/stretchr/testify/assert"
)

func TestOrbitCount(t *testing.T) {
	input := strings.TrimSpace(string(lib.MustReadFile("exinput.txt")))
	got := OrbitCount(input)
	assert.Equal(t, 42, got)
}

func TestSANPath(t *testing.T) {
	input := strings.TrimSpace(string(lib.MustReadFile("exinput2.txt")))
	got := CalcOrbitTransfers(input, "YOU", "SAN")
	assert.Equal(t, 4, got)
}
