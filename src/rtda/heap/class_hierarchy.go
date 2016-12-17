package heap

func (self *Class) IsAssignableFrom(cls *Class) bool {
	return self == cls ||
		self.isSuperClassOf(cls) ||
		self.isSuperInterfaceOf(cls)
}

func (self *Class) isImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

func (self *Class) isSuperClassOf(c *Class) bool {
	return c.isSubClassOf(self)
}

func (self *Class) isSuperInterfaceOf(iface *Class) bool {
	return iface.isSubInterfaceOf(self)
}

func (self *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}
