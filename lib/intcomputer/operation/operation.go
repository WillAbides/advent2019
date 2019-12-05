package operation

import (
	"advent2019/lib"
)

type ParameterMode int

type Computer interface {
	NextPtr() int
	NextInt() int
	WritePosition(pos, val int)
	NextInput() int
	Output(int)
	SetCursor(int)
	Stop()
}

const (
	PositionMode  ParameterMode = 0
	ImmediateMode ParameterMode = 1
)

const NoOp = Operation(-999)

type Operation int

func (o Operation) OpCode() int {
	return int(o) % 100
}

func (o Operation) ParamMode(param int) ParameterMode {
	dgts := lib.ReverseInts(lib.IntDigits(int(o)))
	if len(dgts) <= param+2 {
		return PositionMode
	}
	switch dgts[param+2] {
	case 1:
		return ImmediateMode
	default:
		return PositionMode
	}
}

func (o Operation) ParamValues(c Computer, count int) []int {
	result := make([]int, count)
	for i := 0; i < count; i++ {
		switch o.ParamMode(i) {
		case PositionMode:
			result[i] = c.NextPtr()
		case ImmediateMode:
			result[i] = c.NextInt()
		default:
			panic("unknown mode")
		}
	}
	return result
}

type OpFunc func(op Operation, c Computer)

