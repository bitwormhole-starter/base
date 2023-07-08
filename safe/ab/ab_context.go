package ab

import "github.com/starter-go/base/safe"

// context 是一个双模式 (A-B) 上下文
type context struct {
	modeA      safe.Mode
	modeB      safe.Mode
	modeFacade safe.Mode
	usb        bool // use mode b
}

func (inst *context) getLock(l *facadeLock) safe.Lock {
	if inst.usb {
		return l.lb
	}
	return l.la
}

////////////////////////////////////////////////////////////////////////////////
