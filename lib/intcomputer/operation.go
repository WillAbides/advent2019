package intcomputer

import (
	"github.com/WillAbides/advent2019/lib"
)

type ParameterMode int64

type Computer interface {
	NextPtr(rel bool) (int64, error)
	NextInt() (int64, error)
	WritePosition(pos int64, rel bool, val int64) error
	NextInput() (int64, error)
	Output(int64) error
	SetCursor(int64) error
	UpdateRelativeBase(int64)
	Stop()
	NewError(string) *ComputerErr
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

func (o Operation) ParamValues(c Computer, count int64) ([]int64, error) {
	result := make([]int64, count)
	var err error
	for i := int64(0); i < count; i++ {
		switch o.ParamMode(i) {
		case PositionMode:
			result[i], err = c.NextPtr(false)
		case RelativeMode:
			result[i], err = c.NextPtr(true)
		case ImmediateMode:
			result[i], err = c.NextInt()
		default:
			return result, c.NewError("unknown mode")
		}
	}
	return result, err
}

type OpFunc func(op Operation, c Computer) error

var OpFuncs = map[int64]OpFunc{
	// Opcode 1 adds together numbers read from two positions and stores the result in a third position. The three
	//integers immediately after the opcode tell you these three positions - the first two indicate the positions from
	//which you should read the input values, and the third indicates the position at which the output should be stored.
	1: func(op Operation, c Computer) error {
		params, err := op.ParamValues(c, 2)
		if err != nil {
			return err
		}
		rel := false
		if op.ParamMode(int64(len(params))) == RelativeMode {
			rel = true
		}
		nextInt, err := c.NextInt()
		if err != nil {
			return err
		}
		return c.WritePosition(nextInt, rel, params[0]+params[1])
	},

	//Opcode 2 works exactly like opcode 1, except it multiplies the two inputs instead of adding them. Again, the three
	//integers after the opcode indicate
	2: func(op Operation, c Computer) error {
		params, err := op.ParamValues(c, 2)
		if err != nil {
			return err
		}
		rel := false
		if op.ParamMode(int64(len(params))) == RelativeMode {
			rel = true
		}
		nextInt, err := c.NextInt()
		if err != nil {
			return err
		}
		return c.WritePosition(nextInt, rel, params[0]*params[1])
	},

	// Opcode 3 takes a single integer as input and saves it to the position given by its only parameter. For example,
	//the instruction 3,50 would take an input value and store it at address 50.
	3: func(op Operation, c Computer) error {
		rel := false
		if op.ParamMode(0) == RelativeMode {
			rel = true
		}
		nextInt, err := c.NextInt()
		if err != nil {
			return err
		}
		input, err := c.NextInput()
		if err != nil {
			return err
		}
		return c.WritePosition(nextInt, rel, input)
	},

	//Opcode 4 outputs the value of its only parameter. For example, the instruction 4,50 would output the value at
	//address 50.
	4: func(op Operation, c Computer) error {
		params, err := op.ParamValues(c, 1)
		if err != nil {
			return err
		}
		return c.Output(params[0])
	},

	// Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value from
	//the second parameter. Otherwise, it does nothing.
	5: func(op Operation, c Computer) error {
		params, err := op.ParamValues(c, 2)
		if err != nil {
			return err
		}
		if params[0] != 0 {
			err = c.SetCursor(params[1])
		}
		return err
	},

	// Opcode 6 is jump-if-false: if the first parameter is zero, it sets the instruction pointer to the value from the
	//second parameter. Otherwise, it does nothing.
	6: func(op Operation, c Computer) error {
		params, err := op.ParamValues(c, 2)
		if err != nil {
			return err
		}
		if params[0] == 0 {
			err = c.SetCursor(params[1])
		}
		return err
	},

	//Opcode 7 is less than: if the first parameter is less than the second parameter, it stores 1 in the position given by
	//the third parameter. Otherwise, it stores 0.
	7: func(op Operation, c Computer) error {
		params, err := op.ParamValues(c, 2)
		if err != nil {
			return err
		}
		val := int64(0)
		if params[0] < params[1] {
			val = 1
		}
		rel := false
		if op.ParamMode(int64(len(params))) == RelativeMode {
			rel = true
		}
		nextInt, err := c.NextInt()
		if err != nil {
			return err
		}
		return c.WritePosition(nextInt, rel, val)
	},

	//Opcode 8 is equals: if the first parameter is equal to the second parameter, it stores 1 in the position given by the
	//third parameter. Otherwise, it stores 0.
	8: func(op Operation, c Computer) error {
		params, err := op.ParamValues(c, 2)
		if err != nil {
			return err
		}
		val := int64(0)
		if params[0] == params[1] {
			val = 1
		}
		rel := false
		if op.ParamMode(int64(len(params))) == RelativeMode {
			rel = true
		}
		nextInt, err := c.NextInt()
		if err != nil {
			return err
		}
		return c.WritePosition(nextInt, rel, val)
	},

	//Opcode 9 adjusts the relative base by the value of its only parameter. The relative base increases (or decreases,
	//if the value is negative) by the value of the parameter.
	9: func(op Operation, c Computer) error {
		params, err := op.ParamValues(c, 1)
		if err != nil {
			return err
		}
		c.UpdateRelativeBase(params[0])
		return nil
	},

	//Opcode 99 stops processing operations
	99: func(op Operation, c Computer) error {
		c.Stop()
		return nil
	},
}
