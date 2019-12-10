package day7

import (
	"testing"

	"github.com/WillAbides/advent2019/lib"
	"github.com/WillAbides/advent2019/lib/amplifier"

	"github.com/stretchr/testify/assert"
)

func getFinalOutput(amp amplifier.Amplifier) int64 {
	amp.Start()
	amp.Input(0)
	amp.Wait()
	return amp.LastOutput()
}

func series(phaseSettings, program []int64) amplifier.Amplifier {
	amps := make([]amplifier.Amplifier, len(phaseSettings))
	for i, setting := range phaseSettings {
		amps[i] = amplifier.New(setting, program)
	}
	amp := amplifier.NewSeriesAmp(amps)
	amplifier.ConnectAmps(amp, amp)
	return amp
}

func maxSignal(phases, program []int64) int64 {
	var maxSig int64
	maxPhases := make([]int64, len(phases))
	for _, ps := range lib.IntSlicePermutations(phases) {
		amp := series(ps, program)
		sig := getFinalOutput(amp)
		if sig > maxSig {
			copy(maxPhases, ps)
			maxSig = sig
		}
	}
	return maxSig
}

func TestPart1(t *testing.T) {
	input := string(lib.MustReadFile("input.txt"))
	program := lib.CSInts(input)
	got := maxSignal([]int64{0, 1, 2, 3, 4}, program)
	assert.Equal(t, int64(366376), got)
}

func TestPart2(t *testing.T) {
	input := string(lib.MustReadFile("input.txt"))
	program := lib.CSInts(input)
	got := maxSignal([]int64{5, 6, 7, 8, 9}, program)
	assert.Equal(t, int64(21596786), got)
}
