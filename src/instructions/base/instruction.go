package base

import "rtda"

type Instrucion interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}


type NoOperandInstruction struct {}

type BranchInstruction struct {
	Offset int
}

type Index8Instruction struct {
	Index int
}

type Index16Instruction struct {
	Index uint
}

func (self *NoOperandInstruction) FetchOperands(reader *BytecodeReader)  {
	//?
}



func (self *BranchInstruction) FetchOperands(reader *BytecodeReader)  {
	self.Offset = int(reader.ReadInt16())
}



func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}


func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
