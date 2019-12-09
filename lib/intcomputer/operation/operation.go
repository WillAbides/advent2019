package operation

import (
	"github.com/WillAbides/advent2019/lib"
)

type ParameterMode int64

type Computer interface {
	NextPtr(rel bool) int64
	NextInt() int64
	WritePosition(pos int64, rel bool, val int64)
	NextInput() int64
	Output(int64)
	SetCursor(int64)
	UpdateRelativeBase(int64)
	Stop()
}

const (
	PositionMode  ParameterMode = 0
	ImmediateMode ParameterMode = 1
	RelativeMode  ParameterMode = 2
)

const NoOp = Operation(-999)

type Operation int

func (o Operation) OpCode() int64 {
	return int64(o) % 100
}

func (o Operation) ParamMode(param int64) ParameterMode {
	dgts := lib.ReverseInts(lib.IntDigits(int(o)))
	if int64(len(dgts)) <= param+2 {
		return PositionMode
	}
	switch dgts[param+2] {
	case 1:
		return ImmediateMode
	case 2:
		return RelativeMode
	default:
		return PositionMode
	}
}

func (o Operation) ParamValues(c Computer, count int64) []int64 {
	result := make([]int64, count)
	for i := int64(0); i < count; i++ {
		switch o.ParamMode(i) {
		case PositionMode:
			result[i] = c.NextPtr(false)
		case RelativeMode:
			result[i] = c.NextPtr(true)
		case ImmediateMode:
			result[i] = c.NextInt()
		default:
			panic("unknown mode")
		}
	}
	return result
}

type OpFunc func(op Operation, c Computer)
