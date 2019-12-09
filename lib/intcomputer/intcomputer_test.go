package intcomputer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutput(t *testing.T) {
	for _, td := range []struct {
		codes  []int64
		inputs []int64
		want   []int64
	}{
		{
			codes: []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
			want:  []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		},
		{
			codes: []int64{109, 1, 204, -1, 1001, 0, 0, 0, 99},
			want:  []int64{109},
		},
		{
			codes: []int64{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
			want:  []int64{1219070632396864},
		},
		{
			codes: []int64{104, 1125899906842624, 99},
			want:  []int64{1125899906842624},
		},
		{
			codes:  []int64{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			inputs: []int64{8},
			want:   []int64{1},
		},
		{
			codes:  []int64{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			inputs: []int64{7},
			want:   []int64{0},
		},
		{
			codes:  []int64{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			inputs: []int64{7},
			want:   []int64{1},
		},
		{
			codes:  []int64{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			inputs: []int64{8},
			want:   []int64{0},
		},
		{
			codes:  []int64{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			inputs: []int64{8},
			want:   []int64{1},
		},
		{
			codes:  []int64{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			inputs: []int64{7},
			want:   []int64{0},
		},
		{
			codes:  []int64{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			inputs: []int64{7},
			want:   []int64{1},
		},
		{
			codes:  []int64{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			inputs: []int64{8},
			want:   []int64{0},
		},
		{
			codes:  []int64{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			inputs: []int64{0},
			want:   []int64{0},
		},
		{
			codes:  []int64{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			inputs: []int64{6},
			want:   []int64{1},
		},
		{
			codes:  []int64{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			inputs: []int64{0},
			want:   []int64{0},
		},
		{
			codes:  []int64{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			inputs: []int64{6},
			want:   []int64{1},
		},
		{
			codes: []int64{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			inputs: []int64{7},
			want:   []int64{999},
		},
		{
			codes: []int64{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			inputs: []int64{8},
			want:   []int64{1000},
		},
		{
			codes: []int64{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			inputs: []int64{9},
			want:   []int64{1001},
		},
	} {
		t.Run("", func(t *testing.T) {
			output := &OutputRecorder{}
			c := NewIntComputer(td.codes, output.HandleOutput, SimpleInputter(td.inputs...))
			err := c.RunOperations()
			assert.NoError(t, err)
			assert.Equal(t, td.want, output.Outputs)
		})
	}
}

func TestOperation_OpCode(t *testing.T) {
	for _, td := range []struct {
		input Operation
		want  int64
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
	c := NewIntComputer([]int64{1, 2, 3, 4, 5}, nil, nil)
	o := Operation(10101)
	got := o.ParamValues(c.opComputer(), 3)
	assert.Equal(t, []int64{1, 3, 3}, got)
}

func TestOperation_ParamMode(t *testing.T) {
	for _, td := range []struct {
		input    Operation
		position int64
		want     ParameterMode
	}{
		{0, 0, PositionMode},
		{0, 1, PositionMode},
		{101, 0, ImmediateMode},
		{1001, 1, ImmediateMode},
		{10001, 2, ImmediateMode},
		{10001, 1, PositionMode},
	} {
		t.Run(fmt.Sprintf("%d", td.input), func(t *testing.T) {
			assert.Equal(t, td.want, td.input.ParamMode(td.position))
		})
	}
}

func TestIntComputer_RunOperations(t *testing.T) {
	for _, td := range []struct {
		input []int64
		want  []int64
	}{
		{
			[]int64{1, 4, 4, 0, 99},
			[]int64{198, 4, 4, 0, 99},
		},
		{
			[]int64{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			[]int64{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		},
		{
			[]int64{1, 0, 0, 0, 99},
			[]int64{2, 0, 0, 0, 99},
		},
		{
			[]int64{2, 3, 0, 3, 99},
			[]int64{2, 3, 0, 6, 99},
		}, {
			[]int64{2, 4, 4, 5, 99, 0},
			[]int64{2, 4, 4, 5, 99, 9801},
		}, {
			[]int64{1, 1, 1, 4, 99, 5, 6, 0, 99},
			[]int64{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	} {
		t.Run("", func(t *testing.T) {
			makeTestMemory := func(input []int64) map[int64]int64 {
				mem := make(map[int64]int64, len(input))
				for k, v := range input {
					mem[int64(k)] = v
				}
				return mem
			}
			c := NewIntComputer(td.input, nil, nil)
			assert.NoError(t, c.RunOperations())
			wantMem := makeTestMemory(td.want)
			assert.Equal(t, wantMem, c.memory)
		})
	}
}
