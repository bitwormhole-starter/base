package util

import "sort"

// LenFn 定义 Len 函数
type LenFn func() int

// LessFn 定义 Less 函数
type LessFn func(i1, i2 int) bool

// SwapFn 定义 Swap 函数
type SwapFn func(i1, i2 int)

// Sorter 是一个排序工具
type Sorter struct {
	OnLen  LenFn
	OnLess LessFn
	OnSwap SwapFn
}

func (inst *Sorter) Len() int {
	return inst.OnLen()
}

func (inst *Sorter) Less(i1, i2 int) bool {
	return inst.OnLess(i1, i2)
}

func (inst *Sorter) Swap(i1, i2 int) {
	inst.OnSwap(i1, i2)
}

// Sort 排序
func (inst *Sorter) Sort() {
	sort.Sort(inst)
}
