package rtda

type Frame struct {
	lower		*Frame
	localValue	LocalVars
	operandStack	*OperandStack
}

func NewFrame(maxLocals,maxStack uint) *Frame {
	return &Frame{
		localVars:	newLocalVars(maxLocals),
		operandStack:	newOperandStack(maxStack),
	}
}


