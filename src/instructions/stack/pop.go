package stack

import (
	"instructions/base"
	"rtda"
)

type POP struct{ base.NoOperandsInstruction }

type POP2 struct{ base.NoOperandsInstruction }

//弹出int、float等占用一个位置的变量
func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

//两个
func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
