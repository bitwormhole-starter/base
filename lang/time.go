package lang

import "time"

// Time 用 int64 表示一个 unix 时间戳，基于 UTC-1970-01-01
type Time int64

func (t Time) String() string {
	tt := t.Time()
	return tt.String()
}

// Time 把时间戳转换为 time.Time 形式
func (t Time) Time() time.Time {
	ms := int64(t)
	return time.UnixMilli(ms)
}

// Int 把时间戳转换为 int64 形式
func (t Time) Int() int64 {
	return int64(t)
}

// NewTime 根据 time.Time 创建时间戳
func NewTime(t time.Time) Time {
	ms := t.UnixMilli()
	return Time(ms)
}

// Now 取当前时间戳
func Now() Time {
	now := time.Now()
	return NewTime(now)
}
