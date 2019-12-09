package intcomputer

import (
	"fmt"
	"testing"

	"advent2019/lib/intcomputer/operation"

	"github.com/stretchr/testify/assert"
)

func TestOutput(t *testing.T) {
	for _, td := range []struct {
		codes  []int
		inputs []int
		want   []int
	}{
		{
			codes: []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
			want:  []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		},
		{
			codes: []int{109, 1, 204, -1, 1001},
			want:  []int{109},
		},
		{
			codes: []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
			want:  []int{1219070632396864},
		},
		{
			codes: []int{104, 1125899906842624, 99},
			want:  []int{1125899906842624},
		},
		{
			codes:  []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			inputs: []int{8},
			want:   []int{1},
		},
		{
			codes:  []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			inputs: []int{7},
			want:   []int{0},
		},
		{
			codes:  []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			inputs: []int{7},
			want:   []int{1},
		},
		{
			codes:  []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			inputs: []int{8},
			want:   []int{0},
		},
		{
			codes:  []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			inputs: []int{8},
			want:   []int{1},
		},
		{
			codes:  []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			inputs: []int{7},
			want:   []int{0},
		},
		{
			codes:  []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			inputs: []int{7},
			want:   []int{1},
		},
		{
			codes:  []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			inputs: []int{8},
			want:   []int{0},
		},
		{
			codes:  []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			inputs: []int{0},
			want:   []int{0},
		},
		{
			codes:  []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			inputs: []int{6},
			want:   []int{1},
		},
		{
			codes:  []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			inputs: []int{0},
			want:   []int{0},
		},
		{
			codes:  []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			inputs: []int{6},
			want:   []int{1},
		},
		{
			codes: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			inputs: []int{7},
			want:   []int{999},
		},
		{
			codes: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			inputs: []int{8},
			want:   []int{1000},
		},
		{
			codes: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			inputs: []int{9},
			want:   []int{1001},
		},
	} {
		t.Run("", func(t *testing.T) {
			output := &OutputRecorder{}
			c := NewIntComputer(td.codes, output.HandleOutput, SimpleInputter(td.inputs...))
			c.RunOperations()
			assert.Equal(t, td.want, output.Outputs)
		})
	}
}

func TestOperation_OpCode(t *testing.T) {
	for _, td := range []struct {
		input operation.Operation
		want  int
	}{
		{0, 0},
		{1, 1},
		{10, 10},
		{1010, 10},
	} {
		t.Run(fmt.Sprintf("%d", td.input), func(t *testing.T) {
			assert.Equal(t, td.want, td.input.OpCode())
		})
	}
}

func TestOperation_ParamValues(t *testing.T) {
	c := NewIntComputer([]int{1, 2, 3, 4, 5}, nil, nil)
	o := operation.Operation(10101)
	got := o.ParamValues(c.opComputer(), 3)
	assert.Equal(t, []int{1, 3, 3}, got)
}

func TestOperation_ParamMode(t *testing.T) {
	for _, td := range []struct {
		input    operation.Operation
		position int
		want     operation.ParameterMode
	}{
		{0, 0, operation.PositionMode},
		{0, 1, operation.PositionMode},
		{101, 0, operation.ImmediateMode},
		{1001, 1, operation.ImmediateMode},
		{10001, 2, operation.ImmediateMode},
		{10001, 1, operation.PositionMode},
	} {
		t.Run(fmt.Sprintf("%d", td.input), func(t *testing.T) {
			assert.Equal(t, td.want, td.input.ParamMode(td.position))
		})
	}
}

func TestIntComputer_RunOperations(t *testing.T) {
	for _, td := range []struct {
		input []int
		want  []int
	}{
		{
			[]int{1, 4, 4, 0, 99},
			[]int{198, 4, 4, 0, 99},
		},
		{
			[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			[]int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		},
		{
			[]int{1, 0, 0, 0, 99},
			[]int{2, 0, 0, 0, 99},
		},
		{
			[]int{2, 3, 0, 3, 99},
			[]int{2, 3, 0, 6, 99},
		}, {
			[]int{2, 4, 4, 5, 99, 0},
			[]int{2, 4, 4, 5, 99, 9801},
		}, {
			[]int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			[]int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	} {
		t.Run("", func(t *testing.T) {
			makeTestMemory := func(input []int) map[int]int {
				mem := make(map[int]int, len(input))
				for k, v := range input {
					mem[k] = v
				}
				return mem
			}
			c := NewIntComputer(td.input, nil, nil)
			c.RunOperations()
			wantMem := makeTestMemory(td.want)
			assert.Equal(t, wantMem, c.memory)
		})
	}
}
