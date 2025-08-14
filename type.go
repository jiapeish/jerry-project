package main

type IntensitySegments struct {
	diff  map[int]int // k->v: 从区间跳变点到Intensity的delta的映射，这里v是Intensity的变化量delta，作用于[k, +Inf)
	keys  []int       // 排好序的区间跳变点
	cache string      // 当前的segment pairs: 每pair包括start point和intensity
}

func NewIntensitySegments() *IntensitySegments {
	return &IntensitySegments{
		diff:  make(map[int]int),
		keys:  make([]int, 0),
		cache: "[]",
	}
}
