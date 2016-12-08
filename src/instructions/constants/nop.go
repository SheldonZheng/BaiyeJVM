package constants

import "instructions/base"
import "rtda"

type NOP struct {
	base.NoOperandInstruction
}

func (self *NOP) Execute(frame *rtda.Frame)  {
	//?
}
