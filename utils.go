package main

import (
	"sort"
)

// intensityAt 返回x处的intensity (其左侧的delta会对其有影响).
func (is *IntensitySegments) intensityAt(x int) int {
	sum := 0
	for _, k := range is.keys {
		// 遇到大于 x 的断点，就停止遍历，因为右侧的点对其intensity没有影响哈
		if k > x {
			break
		}
		sum += is.diff[k]
	}
	return sum
}

// putDelta 修改k的intensity
func (is *IntensitySegments) putDelta(k, delta int) {
	if delta == 0 {
		return
	}

	// 如果跳变点k已经存在
	if v, ok := is.diff[k]; ok {
		v += delta
		if v == 0 {
			// 累加后delta变为0，没有用了，需要删除
			delete(is.diff, k)
			idx := lowerBound(is.keys, k)
			if idx < len(is.keys) && is.keys[idx] == k {
				is.keys = append(is.keys[:idx], is.keys[idx+1:]...)
			}
			return
		}
		is.diff[k] = v
		return
	}

	// 新的跳变点
	is.diff[k] = delta
	idx := lowerBound(is.keys, k)
	// 根据位置插入k到数组keys
	if idx == len(is.keys) {
		is.keys = append(is.keys, k)
	} else if is.keys[idx] != k {
		is.keys = append(is.keys, 0)
		copy(is.keys[idx+1:], is.keys[idx:])
		is.keys[idx] = k
	}
}

func lowerBound(a []int, x int) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

func upperBound(a []int, x int) int {
	return sort.Search(len(a), func(i int) bool { return a[i] > x })
}

// 丢弃数组中的区间[l, r]
func dropout(a []int, l, r int) []int {
	if l >= r {
		return a
	}
	out := make([]int, 0, len(a)-(r-l))
	out = append(out, a[:l]...)
	out = append(out, a[r:]...)
	return out
}
