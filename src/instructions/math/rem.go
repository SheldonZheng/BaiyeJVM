package math

import (
	"instructions/base"
	"math"
	"rtda"
)

// Double 求余
type DREM struct{ base.NoOperandsInstruction }

func (self *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	//TODO
	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}

// Float 求余
type FREM struct{ base.NoOperandsInstruction }

func (self *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	//TODO
	result := float32(math.Mod(float64(v1), float64(v2)))
	stack.PushFloat(result)
}

// Int 求余
type IREM struct{ base.NoOperandsInstruction }

func (self *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	} else {
		result := v1 % v2
		stack.PushInt(result)
	}
}

// Long 求余
type LREM struct{ base.NoOperandsInstruction }

func (self *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()

	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	} else {
		result := v1 % v2
		stack.PushLong(result)
	}
}
