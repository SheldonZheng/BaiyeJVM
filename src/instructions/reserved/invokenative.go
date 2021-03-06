package reserved

import "instructions/base"
import "native"
import "rtda"

import _ "native/java/lang"
import _ "native/java/io"
import _ "native/java/security"
import _ "native/java/util/concurrent/atomic"
import _ "native/sun/misc"
import _ "native/sun/io"
import _ "native/sun/reflect"

type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()
	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}
	nativeMethod(frame)
}
