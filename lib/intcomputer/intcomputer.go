package intcomputer

import (
	"fmt"
)

type ComputerErr struct {
	State   *ComputerState
	Message string
}

func (c *ComputerErr) Error() string {
	return c.Message
}

func SimpleInputter(vals ...int64) func() (int64, error) {
	return func() (int64, error) {
		var val int64
		val, vals = vals[len(vals)-1], vals[:len(vals)-1]
		return val, nil
	}
}

type OutputRecorder struct {
	Outputs []int64
}

func (o *OutputRecorder) HandleOutput(_ *IntComputer, n int64) error {
	o.Outputs = append(o.Outputs, n)
	return nil
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

func (o *opComputer) NextInput() (int64, error) {
	return o.inputter()
}

func (o *opComputer) NextInt() (int64, error) {
	return o.nextInt()
}

func (o *opComputer) NextPtr(rel bool) (int64, error) {
	return o.nextPtr(rel)
}

func (o *opComputer) WritePosition(pos int64, rel bool, val int64) error {
	return o.writePosition(pos, rel, val)
}

func (o *opComputer) SetCursor(n int64) error {
	return o.setCursor(n)
}

func (o *opComputer) UpdateRelativeBase(n int64) {
	o.setRelativeBase(o.relativeBase + n)
}

func (o *opComputer) Output(n int64) error {
	return o.output(n)
}

func (o *opComputer) NewError(message string) *ComputerErr {
	return o.Computer().NewError(message)
}

var _ Computer = &opComputer{}

type Inputter func() (int64, error)

type OutputHandler func(c *IntComputer, n int64) error

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

type ComputerState struct {
	Memory       map[int64]int64
	Cursor       int64
	RelativeBase int64
	Stopped      bool
}

func (c *ComputerState) String() string {
	return fmt.Sprintf("%#v", c)
}

type IntComputer struct {
	relativeBase  int64
	memory        map[int64]int64
	cursor        int64
	opFuncs       map[int64]OpFunc
	stopped       bool
	inputter      Inputter
	outputHandler OutputHandler
	onStop        func(c *IntComputer)
	_opComputer   *opComputer
}

func (c *IntComputer) NewError(message string) *ComputerErr {
	return &ComputerErr{
		Message: message,
		State:   c.State(),
	}
}

func (c *IntComputer) SetOnStop(onStop func(c *IntComputer)) {
	c.onStop = onStop
}

func (c *IntComputer) State() *ComputerState {
	mem := make(map[int64]int64, len(c.memory))
	for k, v := range c.memory {
		mem[k] = v
	}
	return &ComputerState{
		Memory:       mem,
		Cursor:       c.cursor,
		RelativeBase: c.relativeBase,
		Stopped:      c.stopped,
	}
}

func (c *IntComputer) output(n int64) error {
	if c.outputHandler != nil {
		return c.outputHandler(c, n)
	}
	fmt.Println(n)
	return nil
}

func (c *IntComputer) opComputer() *opComputer {
	if c._opComputer == nil {
		c._opComputer = &opComputer{c}
	}
	return c._opComputer
}

func (c *IntComputer) RunOperations() error {
	if c.opFuncs == nil {
		c.opFuncs = OpFuncs
	}
	for {
		if c.stopped {
			return nil
		}
		op, err := c.nextOperation()
		if err != nil {
			return err
		}
		if op == NoOp {
			return &ComputerErr{
				Message: "encountered unexpected NoOp",
				State:   c.State(),
			}
		}
		opFunc := c.opFuncs[op.OpCode()]
		if opFunc == nil {
			return &ComputerErr{
				Message: "encountered unknown operation",
				State:   c.State(),
			}
		}
		err = opFunc(op, c.opComputer())
		if err != nil {
			return err
		}
	}
}

func (c *IntComputer) setCursor(n int64) error {
	if n < 0 {
		return c.NewError("attempting to set cursor to negative value")
	}
	c.cursor = n
	return nil
}

func (c *IntComputer) setRelativeBase(n int64) {
	c.relativeBase = n
}

func (c *IntComputer) nextOperation() (Operation, error) {
	_, ok := c.memory[c.cursor]
	if !ok {
		return NoOp, nil
	}
	nextInt, err := c.nextInt()
	if err != nil {
		return NoOp, err
	}
	return Operation(nextInt), nil
}

func (c *IntComputer) nasNext() bool {
	_, ok := c.memory[c.cursor+1]
	return ok
}

func (c *IntComputer) ReadPosition(pos int64) (int64, error) {
	if pos < 0 {
		return 0, c.NewError("tried reading from a negative position")
	}
	return c.memory[pos], nil
}

func (c *IntComputer) writePosition(pos int64, rel bool, val int64) error {
	target := pos
	if rel {
		target = pos + c.relativeBase
	}
	if target < 0 {
		return c.NewError("attempting to write to a negative register")
	}
	c.memory[target] = val
	return nil
}

func (c *IntComputer) nextInt() (int64, error) {
	c.cursor++
	return c.ReadPosition(c.cursor - 1)
}

func (c *IntComputer) nextPtr(rel bool) (int64, error) {
	ptr, err := c.nextInt()
	if err != nil {
		return 0, err
	}
	if rel {
		ptr = ptr + c.relativeBase
	}
	return c.ReadPosition(ptr)
}

func (c *IntComputer) stop() {
	if c.onStop != nil {
		c.onStop(c)
	}
	c.stopped = true
}
