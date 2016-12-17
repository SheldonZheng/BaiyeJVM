package heap

import (
	. "classfile"
	"sync"
)


type Class struct {
	accessFlags uint16
	name string // thisClassName
	superClassName string
	interfaceNames []string
	constantPool *ConstantPool
	fields []*Field
	methods []*Method
	loader *ClassLoader
	superClass *Class
	interfaces []*Class
	instanceSlotCount uint
	staticSlotCount uint
	staticVars *Slots
}

func newClass(cf *ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}

func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}
func (self *Class) Name() string {
	return self.name
}
func (self *Class) Methods() []*Method {
	return self.methods
}
func (self *Class) Fields() []*Field {
	return self.fields
}
func (self *Class) SuperClass() *Class {
	return self.superClass
}
func (self *Class) Interfaces() []*Class {
	return self.interfaces
}