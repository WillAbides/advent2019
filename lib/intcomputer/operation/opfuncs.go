package operation

var OpFuncs = map[int]OpFunc{
	// Opcode 1 adds together numbers read from two positions and stores the result in a third position. The three
	//integers immediately after the opcode tell you these three positions - the first two indicate the positions from
	//which you should read the input values, and the third indicates the position at which the output should be stored.
	1: func(op Operation, c Computer) {
		params := op.ParamValues(c, 2)
		c.WritePosition(c.NextInt(), params[0]+params[1])
	},

	//Opcode 2 works exactly like opcode 1, except it multiplies the two inputs instead of adding them. Again, the three
	//integers after the opcode indicate
	2: func(op Operation, c Computer) {
		params := op.ParamValues(c, 2)
		c.WritePosition(c.NextInt(), params[0]*params[1])
	},

	// Opcode 3 takes a single integer as input and saves it to the position given by its only parameter. For example,
	//the instruction 3,50 would take an input value and store it at address 50.
	3: func(op Operation, c Computer) {
		c.WritePosition(c.NextInt(), c.NextInput())
	},

	//Opcode 4 outputs the value of its only parameter. For example, the instruction 4,50 would output the value at
	//address 50.
	4: func(op Operation, c Computer) {
		params := op.ParamValues(c, 1)
		c.Output(params[0])
	},

	// Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value from
	//the second parameter. Otherwise, it does nothing.
	5: func(op Operation, c Computer) {
		params := op.ParamValues(c, 2)
		if params[0] != 0 {
			c.SetCursor(params[1])
		}
	},

	// Opcode 6 is jump-if-false: if the first parameter is zero, it sets the instruction pointer to the value from the
	//second parameter. Otherwise, it does nothing.
	6: func(op Operation, c Computer) {
		params := op.ParamValues(c, 2)
		if params[0] == 0 {
			c.SetCursor(params[1])
		}
	},

	//Opcode 7 is less than: if the first parameter is less than the second parameter, it stores 1 in the position given by
	//the third parameter. Otherwise, it stores 0.
	7: func(op Operation, c Computer) {
		params := op.ParamValues(c, 2)
		val := 0
		if params[0] < params[1] {
			val = 1
		}
		c.WritePosition(c.NextInt(), val)
	},

	//Opcode 8 is equals: if the first parameter is equal to the second parameter, it stores 1 in the position given by the
	//third parameter. Otherwise, it stores 0.
	8: func(op Operation, c Computer) {
		params := op.ParamValues(c, 2)
		val := 0
		if params[0] == params[1] {
			val = 1
		}
		c.WritePosition(c.NextInt(), val)
	},

	//Opcode 99 stops processing operations
	99: func(op Operation, c Computer) {
		c.Stop()
	},
}
