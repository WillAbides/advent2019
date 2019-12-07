package day7

import (
	"sync"
	"testing"
	"time"

	"advent2019/lib"
	"advent2019/lib/intcomputer"

	"github.com/stretchr/testify/assert"
)

type amp struct {
	phaseSetting int
	program      []int
	inputCh      chan int
	outputCh     chan int
	lastOutput   int
}

func (a *amp) writeInput(input int) {
	a.inputCh <- input
}

func (a *amp) start(done func()) {
	inputter := func() int {
		got := <-a.inputCh
		return got
	}
	outputHandler := func(c *intcomputer.IntComputer, n int) {
		a.lastOutput = n
		select {
		case a.outputCh <- n:
		case <-time.After(time.Millisecond):
		}
	}
	computer := intcomputer.NewIntComputer(a.program, outputHandler, inputter)
	computer.SetOnStop(func(_ *intcomputer.IntComputer) {
		done()
	})
	go computer.RunOperations()
	a.inputCh <- a.phaseSetting
}

type amps []*amp

func (a amps) run() int {
	var wg sync.WaitGroup
	for i := range a {
		j := i
		wg.Add(1)
		done := func() {
			//fmt.Println(j)
			wg.Done()
		}
		a[j].start(done)
	}
	a[0].inputCh <- 0
	//wg.Wait()
	time.Sleep(2 * time.Millisecond)
	output := a[len(a)-1].lastOutput
	return output
}

func newAmps(phaseSettings, program []int) amps {
	na := make(amps, len(phaseSettings))
	for i, setting := range phaseSettings {
		pr := make([]int, len(program))
		copy(pr, program)
		var in chan int
		if i > 0 {
			in = na[i-1].outputCh
		}
		na[i] = &amp{
			program:      pr,
			phaseSetting: setting,
			inputCh:      in,
			outputCh:     make(chan int),
		}
	}
	na[0].inputCh = na[len(na)-1].outputCh
	return na
}

func maxSignal(phases, program []int) int {
	maxSig := 0
	maxPhases := make([]int, len(phases))
	for _, ps := range lib.IntSlicePermutations(phases) {
		sig := newAmps(ps, program).run()
		if sig > maxSig {
			copy(maxPhases, ps)
			maxSig = sig
		}
	}
	return maxSig
}

func TestAmps(t *testing.T) {
	tests := []struct {
		program       []int
		phaseSettings []int
		want          int
	}{
		{
			program:       []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
			phaseSettings: []int{4, 3, 2, 1, 0},
			want:          43210,
		},
		{
			program:       []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
			phaseSettings: []int{0, 1, 2, 3, 4},
			want:          54321,
		},
		{
			program: []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33,
				1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
			phaseSettings: []int{1, 0, 4, 3, 2},
			want:          65210,
		},
		{
			program: []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26,
				27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5},
			phaseSettings: []int{9, 8, 7, 6, 5},
			want:          139629729,
		},
		{
			program: []int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54,
				-5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4,
				53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10},
			phaseSettings: []int{9, 7, 8, 5, 6},
			want:          18216,
		},
		{
			program:       realProgram,
			phaseSettings: []int{2, 3, 0, 4, 1},
			want:          366376,
		},
		{
			program:       realProgram,
			phaseSettings: []int{9, 5, 8, 6, 7},
			want:          21596786,
		},
	}
	for _, td := range tests {
		t.Run("", func(t *testing.T) {
			got := newAmps(td.phaseSettings, td.program).run()
			assert.Equal(t, td.want, got)
		})
	}
}

var realProgram = lib.CSInts(string(lib.MustReadFile("input.txt")))

func Test_maxSignal(t *testing.T) {
	tests := []struct {
		program []int
		phases  []int
		want    int
	}{
		{
			program: realProgram,
			phases:  []int{0, 1, 2, 3, 4},
			want:    366376,
		},
		{
			program: realProgram,
			phases:  []int{5, 6, 7, 8, 9},
			want:    21596786,
		},
		{
			program: []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
			phases:  []int{0, 1, 2, 3, 4},
			want:    43210,
		},
		{
			program: []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
			phases:  []int{0, 1, 2, 3, 4},
			want:    54321,
		},
		{
			program: []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33,
				1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
			phases: []int{0, 1, 2, 3, 4},
			want:   65210,
		},
		{
			program: []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26,
				27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5},
			phases: []int{5, 6, 7, 8, 9},
			want:   139629729,
		},
		{
			program: []int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54,
				-5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4,
				53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10},
			phases: []int{5, 6, 7, 8, 9},
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

func TestPart1(t *testing.T) {
	input := string(lib.MustReadFile("input.txt"))
	program := lib.CSInts(input)
	got := maxSignal([]int{0, 1, 2, 3, 4}, program)
	assert.Equal(t, 366376, got)
}

func TestPart2(t *testing.T) {
	input := string(lib.MustReadFile("input.txt"))
	program := lib.CSInts(input)
	got := maxSignal([]int{5, 6, 7, 8, 9}, program)
	assert.Equal(t, 21596786, got)
}