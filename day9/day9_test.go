package day9

import (
	"testing"

	"github.com/WillAbides/advent2019/lib"
	"github.com/WillAbides/advent2019/lib/intcomputer"

	"github.com/stretchr/testify/assert"
)

var realProgram = lib.CSInts(string(lib.MustReadFile("input.txt")))

func TestPart1(t *testing.T) {
	output := &intcomputer.OutputRecorder{}
	c := intcomputer.NewIntComputer(realProgram, output.HandleOutput, intcomputer.SimpleInputter(1))
	assert.NoError(t, c.RunOperations())
	assert.Equal(t, []int64{3906448201}, output.Outputs)
}

func TestPart2(t *testing.T) {
	output := &intcomputer.OutputRecorder{}
	c := intcomputer.NewIntComputer(realProgram, output.HandleOutput, intcomputer.SimpleInputter(2))
	assert.NoError(t, c.RunOperations())
	assert.Equal(t, []int64{59785}, output.Outputs)
}
