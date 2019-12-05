package intcomputer

import (
	"fmt"

	"advent2019/lib/intcomputer/operation"
)

func SimpleInputter(vals ...int) func() int {
	return func() int {
		var val int
		val, vals = vals[len(vals)-1], vals[:len(vals)-1]
		return val
	}
}

type OutputRecorder struct {
	Outputs []int
}

//opComputer implements operations.Computer
type opComputer struct {
	*IntComputer
}

func (o *opComputer) Stop() {
	o.stop()
}

func (o *opComputer) NextInput() int {
	return o.inputter()
}

func (o *opComputer) NextInt() int {
	return o.nextInt()
}

func (o *opComputer) NextPtr() int {
	return o.nextPtr()
}
func (o *opComputer) WritePosition(pos, val int) {
	o.writePosition(pos, val)
}

func (o *opComputer) SetCursor(n int) {
	o.setCursor(n)
}

func (o *opComputer) Output(n int) {
	o.output(n)
}

var _ operation.Computer = &opComputer{}

func (o *OutputRecorder) HandleOutput(_ *IntComputer, n int) {
	o.Outputs = append(o.Outputs, n)
}

type Inputter func() int

type OutputHandler func(c *IntComputer, n int)

func NewIntComputer(mem []int, outputHandler OutputHandler, inputter Inputter) *IntComputer {
	c := &IntComputer{
		memory:        make([]int, len(mem)),
		inputter:      inputter,
		outputHandler: outputHandler,
	}
	copy(c.memory, mem)
	return c
}

type IntComputer struct {
	memory        []int
	cursor        int
	opFuncs       map[int]operation.OpFunc
	stopped       bool
	inputter      Inputter
	outputHandler OutputHandler
	_opComputer *opComputer
}

func (c *IntComputer) output(n int) {
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

func (c *IntComputer) setCursor(n int) {
	c.cursor = n
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
	return c.cursor < (len(c.memory) - 1)
}

func (c *IntComputer) ReadPosition(pos int) int {
	return c.memory[pos]
}

func (c *IntComputer) writePosition(pos, val int) {
	c.memory[pos] = val
}

func (c *IntComputer) nextInt() int {
	res := c.ReadPosition(c.cursor)
	c.cursor++
	return res
}

func (c *IntComputer) nextPtr() int {
	ptr := c.nextInt()
	return c.ReadPosition(ptr)
}

func (c *IntComputer) stop() {
	c.stopped = true
}
