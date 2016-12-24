package rtda

import "rtda/heap"

type Thread struct {
	pc    int // the address of the instruction currently being executed
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(self, method)
}

func (self *Thread) PC() int {
	return self.pc
}
func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}
func (self *Thread) TopFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) IsStackEmpty() bool {
	return self.stack.isEmpty()
}
