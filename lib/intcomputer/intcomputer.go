package intcomputer

import (
	"fmt"

	"github.com/WillAbides/advent2019/lib/intcomputer/operation"
)

func SimpleInputter(vals ...int64) func() int64 {
	return func() int64 {
		var val int64
		val, vals = vals[len(vals)-1], vals[:len(vals)-1]
		return val
	}
}

type OutputRecorder struct {
	Outputs []int64
}

func (o *OutputRecorder) HandleOutput(_ *IntComputer, n int64) {
	o.Outputs = append(o.Outputs, n)
}

//opComputer implements operations.Computer
type opComputer struct {
	*IntComputer
}

func (o *opComputer) Computer() *IntComputer {
	return o.IntComputer
}

func (o *opComputer) Stop() {
	o.stop()
}

func (o *opComputer) NextInput() int64 {
	return o.inputter()
}

func (o *opComputer) NextInt() int64 {
	return o.nextInt()
}

func (o *opComputer) NextPtr(rel bool) int64 {
	return o.nextPtr(rel)
}

func (o *opComputer) WritePosition(pos int64, rel bool, val int64) {
	o.writePosition(pos, rel, val)
}

func (o *opComputer) SetCursor(n int64) {
	o.setCursor(n)
}

func (o *opComputer) UpdateRelativeBase(n int64) {
	o.setRelativeBase(o.relativeBase + n)
}

func (o *opComputer) Output(n int64) {
	o.output(n)
}

var _ operation.Computer = &opComputer{}

type Inputter func() int64

type OutputHandler func(c *IntComputer, n int64)

func NewIntComputer(mem []int64, outputHandler OutputHandler, inputter Inputter) *IntComputer {
	c := &IntComputer{
		memory:        make(map[int64]int64, len(mem)),
		inputter:      inputter,
		outputHandler: outputHandler,
	}
	for k, v := range mem {
		c.memory[int64(k)] = v
	}
	return c
}

type IntComputer struct {
	relativeBase  int64
	memory        map[int64]int64
	cursor        int64
	opFuncs       map[int64]operation.OpFunc
	stopped       bool
	inputter      Inputter
	outputHandler OutputHandler
	onStop        func(c *IntComputer)
	_opComputer   *opComputer
}

func (c *IntComputer) SetOnStop(onStop func(c *IntComputer)) {
	c.onStop = onStop
}

func (c *IntComputer) output(n int64) {
	if c.outputHandler != nil {
		c.outputHandler(c, n)
		return
	}
	fmt.Println(n)
}

func (c *IntComputer) opComputer() *opComputer {
	if c._opComputer == nil {
		c._opComputer = &opComputer{c}
	}
	return c._opComputer
}

func (c *IntComputer) RunOperations() {
	if c.opFuncs == nil {
		c.opFuncs = operation.OpFuncs
	}
	for {
		if c.stopped {
			return
		}
		op := c.nextOperation()
		if op == operation.NoOp {
			return
		}
		opFunc := c.opFuncs[op.OpCode()]
		if opFunc == nil {
			return
		}
		opFunc(op, c.opComputer())
	}
}

func (c *IntComputer) setCursor(n int64) {
	c.cursor = n
}

func (c *IntComputer) setRelativeBase(n int64) {
	c.relativeBase = n
}

func (c *IntComputer) nextOperation() operation.Operation {
	if !c.nasNext() {
		return operation.NoOp
	}
	return operation.Operation(c.nextInt())
}

func (c *IntComputer) nextOpFunc() operation.OpFunc {
	if !c.nasNext() {
		return nil
	}
	opCode := c.nextInt()
	return c.opFuncs[opCode]
}

func (c *IntComputer) nasNext() bool {
	_, ok := c.memory[c.cursor+1]
	return ok
}

func (c *IntComputer) ReadPosition(pos int64) int64 {
	return c.memory[pos]
}

func (c *IntComputer) writePosition(pos int64, rel bool ,val int64) {
	target := pos
	if rel {
		target = pos + c.relativeBase
	}
	c.memory[target] = val
}

func (c *IntComputer) nextInt() int64 {
	res := c.ReadPosition(c.cursor)
	c.cursor++
	return res
}

func (c *IntComputer) nextPtr(rel bool) int64 {
	ptr := c.nextInt()

	if rel {
		ptr = ptr + c.relativeBase
		_ = ptr
	}
	return c.ReadPosition(ptr)
}

func (c *IntComputer) stop() {
	if c.onStop != nil {
		c.onStop(c)
	}
	c.stopped = true
}
