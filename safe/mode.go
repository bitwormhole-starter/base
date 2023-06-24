package safe

// Mode 表示一种安全模式
type Mode interface {
	NewLock() Lock
}

// Fast 返回快速模式
func Fast() Mode {
	return &modeFast{}
}

// Safe 安全模式
func Safe() Mode {
	return &modeSafe{}
}

// Default 默认模式
func Default() Mode {
	return Safe()
}

////////////////////////////////////////////////////////////////////////////////

type modeFast struct{}

func (inst *modeFast) NewLock() Lock {
	return &lockFast{}
}

////////////////////////////////////////////////////////////////////////////////

type modeSafe struct{}

func (inst *modeSafe) NewLock() Lock {
	return &lockSafe{}
}

////////////////////////////////////////////////////////////////////////////////
