package context2

import (
	"context"
	"fmt"
)

const theValuesKey = "base/context2.Values#binding"

// Getter 是作用于 context.Context 的 key-value 取值接口
type Getter interface {
	GetValue(key any) any
}

// Setter 是作用于 context.Context 的 key-value 设置接口
type Setter interface {
	SetValue(key, value any)
}

// Values 是作用于 context.Context 的 key-value 存取接口
type Values interface {
	Getter
	Setter
	Context() context.Context
}

// GetValues 从上下文中获取 Values 接口
func GetValues(c context.Context) (Values, error) {
	if c == nil {
		return nil, fmt.Errorf("context is nil")
	}
	const key = theValuesKey
	api, ok := c.Value(key).(Values)
	if ok && api != nil {
		return api, nil
	}
	return nil, fmt.Errorf("base/context2: invoking SetupValues() is required before GetValues()")
}

// Setup 安装配置 Values 接口
func Setup(val Values) {
	const key = theValuesKey
	val.SetValue(key, val)
}

// SetupContext 为 context.Context 安装配置 Values 接口
func SetupContext(c context.Context) Values {
	values, _ := GetValues(c)
	if values != nil {
		return values
	}
	if c == nil {
		panic("context is nil")
	}
	v2 := &defaultKeyValues{ctx: c}
	values = v2
	Setup(values)
	return values
}
