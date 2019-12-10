package amplifier

import (
	"testing"

	"github.com/WillAbides/advent2019/lib"

	"github.com/stretchr/testify/assert"
)

var realProgram = lib.CSInts(string(lib.MustReadFile("input.txt")))

func series(phaseSettings, program []int64) Amplifier {
	amps := make([]Amplifier, len(phaseSettings))
	for i, setting := range phaseSettings {
		amps[i] = New(setting, program)
	}
	amp := NewSeriesAmp(amps)
	ConnectAmps(amp, amp)
	return amp
}

func getFinalOutput(amp Amplifier) int64 {
	amp.Start()
	amp.Input(0)
	amp.Wait()
	return amp.LastOutput()
}

func maxSignal(phases, program []int64) int64 {
	var maxSig int64
	maxPhases := make([]int64, len(phases))
	for _, ps := range lib.IntSlicePermutations(phases) {
		amp := series(ps, program)
		sig := getFinalOutput(amp)

		//sig := newAmps(ps, program).run()
		if sig > maxSig {
			copy(maxPhases, ps)
			maxSig = sig
		}
	}
	return maxSig
}

func TestAmps(t *testing.T) {
	tests := []struct {
		program       []int64
		phaseSettings []int64
		want          int64
	}{
		{
			program:       []int64{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
			phaseSettings: []int64{4, 3, 2, 1, 0},
			want:          43210,
		},
		{
			program:       []int64{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
			phaseSettings: []int64{0, 1, 2, 3, 4},
			want:          54321,
		},
		{
			program: []int64{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33,
				1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
			phaseSettings: []int64{1, 0, 4, 3, 2},
			want:          65210,
		},
		{
			program: []int64{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26,
				27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5},
			phaseSettings: []int64{9, 8, 7, 6, 5},
			want:          139629729,
		},
		{
			program: []int64{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54,
				-5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4,
				53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10},
			phaseSettings: []int64{9, 7, 8, 5, 6},
			want:          18216,
		},
		{
			program:       realProgram,
			phaseSettings: []int64{2, 3, 0, 4, 1},
			want:          366376,
		},
		{
			program:       realProgram,
			phaseSettings: []int64{9, 5, 8, 6, 7},
			want:          21596786,
		},
		{
			program:       realProgram,
			phaseSettings: []int64{5, 6, 7, 8, 9},
			want:          7474552,
		},
	}
	for _, td := range tests {
		t.Run("", func(t *testing.T) {
			amp := series(td.phaseSettings, td.program)
			got := getFinalOutput(amp)
			assert.Equal(t, td.want, got)
		})
	}
}

func Test_maxSignal(t *testing.T) {
	tests := []struct {
		program []int64
		phases  []int64
		want    int64
	}{

		{
			program: realProgram,
			phases:  []int64{0, 1, 2, 3, 4},
			want:    366376,
		},
		{
			program: realProgram,
			phases:  []int64{5, 6, 7, 8, 9},
			want:    21596786,
		},
		{
			program: []int64{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
			phases:  []int64{0, 1, 2, 3, 4},
			want:    43210,
		},
		{
			program: []int64{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
			phases:  []int64{0, 1, 2, 3, 4},
			want:    54321,
		},
		{
			program: []int64{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33,
				1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
			phases: []int64{0, 1, 2, 3, 4},
			want:   65210,
		},
		{
			program: []int64{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26,
				27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5},
			phases: []int64{5, 6, 7, 8, 9},
			want:   139629729,
		},
		{
			program: []int64{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54,
				-5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4,
				53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10},
			phases: []int64{5, 6, 7, 8, 9},
			want:   18216,
		},
	}
	for _, td := range tests {
		t.Run("", func(t *testing.T) {
			got := maxSignal(td.phases, td.program)
			assert.Equal(t, td.want, got)
		})
	}
}
