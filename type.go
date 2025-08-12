package main

type IntensitySegments struct {
	diff map[int]int // k->v: 从区间跳变点到Intensity的delta的映射，这里v是Intensity的变化量delta，作用于[k, +Inf)
	keys []int       // 排好序的区间跳变点
}

func NewIntensitySegments() *IntensitySegments {
	return &IntensitySegments{
		diff: make(map[int]int),
		keys: make([]int, 0),
	}
}
