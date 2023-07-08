package ab

import "github.com/starter-go/base/safe"

// Mode 是双状态(A-B)安全模式
type Mode interface {
	safe.Mode

	UseModeA() // 切换到 A 模式
	UseModeB() // 切换到 B 模式
}

////////////////////////////////////////////////////////////////////////////////

// New 新建一个 A-B 模式上下文
func New(a, b safe.Mode) Mode {

	if a == nil {
		a = safe.Default()
	}

	if b == nil {
		b = safe.Default()
	}

	ctx := &context{}
	m := &facadeMode{}

	ctx.modeA = a
	ctx.modeB = b
	ctx.modeFacade = m

	m.context = ctx
	return m
}

////////////////////////////////////////////////////////////////////////////////

type facadeMode struct {
	context *context
}

func (inst *facadeMode) _impl() safe.Mode {
	return inst
}

func (inst *facadeMode) NewLock() safe.Lock {
	ctx := inst.context
	l := &facadeLock{}
	l.context = ctx
	l.la = ctx.modeA.NewLock()
	l.lb = ctx.modeB.NewLock()
	return l
}

// UseModeA 切换到 A 模式
func (inst *facadeMode) UseModeA() {
	inst.context.usb = false
}

// UseModeB 切换到 B 模式
func (inst *facadeMode) UseModeB() {
	inst.context.usb = true
}
