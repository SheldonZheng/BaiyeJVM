package math

import (
	"instructions/base"
	"rtda"
)

//Div Double
type DDIV struct{ base.NoOperandsInstruction }

func (self *DDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 / v2
	stack.PushDouble(result)
}

// Div Float
type FDIV struct{ base.NoOperandsInstruction }

func (self *FDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 / v2
	stack.PushFloat(result)
}

// Div Int
type IDIV struct{ base.NoOperandsInstruction }

func (self *IDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	} else {
		result := v1 / v2
		stack.PushInt(result)
	}
}

//Div Long
type LDIV struct{ base.NoOperandsInstruction }

func (self *LDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()

	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	} else {
		result := v1 / v2
		stack.PushLong(result)
	}
}