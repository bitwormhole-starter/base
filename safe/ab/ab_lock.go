package ab

import "github.com/starter-go/base/safe"

type facadeLock struct {
	context *context
	la      safe.Lock
	lb      safe.Lock
}

func (inst *facadeLock) _impl() safe.Lock {
	return inst
}

func (inst *facadeLock) Lock() {
	l := inst.context.getLock(inst)
	l.Lock()
}

func (inst *facadeLock) Unlock() {
	l := inst.context.getLock(inst)
	l.Unlock()
}
