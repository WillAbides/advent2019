package day9

import (
	"testing"

	"advent2019/lib"
	"advent2019/lib/intcomputer"

	"github.com/stretchr/testify/assert"
)

var realProgram = lib.CSInts(string(lib.MustReadFile("input.txt")))

func TestPart1(t *testing.T) {
	output := &intcomputer.OutputRecorder{}
	c := intcomputer.NewIntComputer(realProgram, output.HandleOutput, intcomputer.SimpleInputter(1))
	c.RunOperations()
	assert.Equal(t, []int{3906448201}, output.Outputs)
}

func TestPart2(t *testing.T) {
	output := &intcomputer.OutputRecorder{}
	c := intcomputer.NewIntComputer(realProgram, output.HandleOutput, intcomputer.SimpleInputter(2))
	c.RunOperations()
	assert.Equal(t, []int{59785}, output.Outputs)
}
