package amplifier

import (
	"fmt"
	"time"

	"github.com/WillAbides/advent2019/lib/intcomputer"
)

type Amplifier interface {
	AddOutputHandler(fn func(int64))
	Wait()
	Start()
	Input(input int64)
	LastOutput() int64
	Running() bool
}

type amplifier struct {
	inputCh       chan int64
	outputHandler func(int64)
	done          chan struct{}
	isDone        bool
	lastOutput    int64
	computer      *intcomputer.IntComputer
}

func (a *amplifier) AddOutputHandler(fn func(int64)) {
	existingHandler := a.outputHandler
	if existingHandler == nil {
		existingHandler = func(int64) {}
	}
	a.outputHandler = func(n int64) {
		fn(n)
		existingHandler(n)
	}
}

func (a *amplifier) Wait() {
	if a.isDone {
		return
	}
	<-a.done
	return
}

func (a *amplifier) output(n int64) {
	a.lastOutput = n
	if a.outputHandler != nil {
		a.outputHandler(n)
	}
}

func New(phaseSetting int64, program []int64) Amplifier {
	programClone := make([]int64, len(program))
	copy(programClone, program)
	amp := &amplifier{
		inputCh: make(chan int64),
		done:    make(chan struct{}),
	}
	outHandler := func(_ *intcomputer.IntComputer, n int64) error {
		amp.output(n)
		return nil
	}
	phaseIsSet := false
	inputter := func() (int64, error) {
		if ! phaseIsSet {
			phaseIsSet = true
			return phaseSetting, nil
		}
		var err error
		var val int64
		select {
		case val = <-amp.inputCh:
		case <-time.After(2 * time.Millisecond):
			fmt.Println("boink")
			err = fmt.Errorf("input timed out")
		}
		return val, err
	}
	amp.computer = intcomputer.NewIntComputer(programClone, outHandler, inputter)
	amp.computer.SetOnStop(func(_ *intcomputer.IntComputer) {
		//fmt.Println("done ", phaseSetting)
		amp.isDone = true
		close(amp.done)
	})

	return amp
}

func (a *amplifier) Start() {
	go a.computer.RunOperations()
}

func (a *amplifier) Input(input int64) {
	a.inputCh <- input
}

func (a *amplifier) LastOutput() int64 {
	return a.lastOutput
}

func (a *amplifier) Running() bool {
	return ! a.isDone
}

type seriesAmp struct {
	amps []Amplifier
}

func NewSeriesAmp(amps []Amplifier) Amplifier {
	for i := range amps {
		if i == 0 {
			continue
		}
		prev := amps[i-1]
		amp := amps[i]
		ConnectAmps(prev, amp)
	}
	return &seriesAmp{
		amps: amps,
	}
}

func (a *seriesAmp) lastAmp() Amplifier {
	return a.amps[len(a.amps)-1]
}

func (a *seriesAmp) AddOutputHandler(fn func(int64)) {
	a.lastAmp().AddOutputHandler(fn)
}

func (a *seriesAmp) Wait() {
	a.lastAmp().Wait()
}

func (a *seriesAmp) Start() {
	for _, amp := range a.amps {
		amp.Start()
	}
}

func (a *seriesAmp) Input(n int64) {
	a.amps[0].Input(n)
}

func (a *seriesAmp) LastOutput() int64 {
	return a.lastAmp().LastOutput()
}

func (a *seriesAmp) Running() bool {
	for _, amp := range a.amps {
		if ! amp.Running() {
			return false
		}
	}
	return true
}

func ConnectAmps(src, dest Amplifier) {
	src.AddOutputHandler(func(n int64) {
		if dest.Running() {
			done := make(chan struct{})
			go func() {
				dest.Input(n)
				close(done)
			}()
			select {
			case <-done:
			case <-time.After(time.Millisecond):
			}
		}
	})
}
