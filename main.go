package main

import (
	"fmt"
	"sort"
	"strings"
)

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

func (is *IntensitySegments) Add(from, to, amount int) {
	if amount == 0 || from >= to {
		return
	}
	is.putDelta(from, amount)
	is.putDelta(to, -amount)
}

// Set 方法把[from, to)的Intensity设置为amount
// 步骤:
//  1. 计算起始点from的调整量: amount - Intensity(from).
//  2. 清除区间(from, to)内的断点并记录其总影响：Set方法会设置[from, to)区间内的Intensity为amount，
//     而区间内的原有跳变点会导致Intensity变化，因此必须删除这些跳变点, 但这些跳变点又会影响to右侧的Intensity
//     直接删除会改变to右侧的Intensity，因此用internal记录这些点的delta和，然后在to处施加这个影响，保证to右侧Intensity不变
//  3. 在from处 +delta，并在to处 -delta+internal
func (is *IntensitySegments) Set(from, to, amount int) {
	if from >= to {
		return
	}
	// Step 1: 计算from处当前的intensity值，从而计算需要调整的量，让from处的强度变为amount
	currentFrom := is.intensityAt(from)
	deltaFrom := amount - currentFrom

	// Step 2: 清除区间内原有的跳变点，但是intensity要记录下来
	left := upperBound(is.keys, from)
	right := lowerBound(is.keys, to)
	internal := 0
	for i := left; i < right; i++ {
		k := is.keys[i]
		internal += is.diff[k]
		delete(is.diff, k)
	}
	// 删除这些跳变点，重合了，没有用了
	is.keys = dropout(is.keys, left, right)

	// Step 3: 调整from和to处的intensity
	is.putDelta(from, deltaFrom)
	is.putDelta(to, internal-deltaFrom)
}

// ToString 打印
func (is *IntensitySegments) ToString() string {
	if len(is.keys) == 0 {
		return "[]"
	}
	var b strings.Builder
	b.WriteByte('[')
	sum := 0
	for i, k := range is.keys {
		sum += is.diff[k]
		if i > 0 {
			b.WriteByte(',')
		}
		fmt.Fprintf(&b, "[%d,%d]", k, sum)
	}
	b.WriteByte(']')
	return b.String()
}

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

func main() {
	// 例子1:
	segments := NewIntensitySegments()
	fmt.Println(segments.ToString()) // "[]"
	segments.Add(10, 30, 1)
	fmt.Println(segments.ToString()) // [[10,1],[30,0]]
	segments.Add(20, 40, 1)
	fmt.Println(segments.ToString()) // [[10,1],[20,2],[30,1],[40,0]]
	segments.Add(10, 40, -2)
	fmt.Println(segments.ToString()) // [[10,-1],[20,0],[30,-1],[40,0]]

	// 例子2:
	segments2 := NewIntensitySegments()
	fmt.Println(segments2.ToString()) // "[]"
	segments2.Add(10, 30, 1)
	fmt.Println(segments2.ToString()) // [[10,1],[30,0]]
	segments2.Add(20, 40, 1)
	fmt.Println(segments2.ToString()) // [[10,1],[20,2],[30,1],[40,0]]
	segments2.Add(10, 40, -1)
	fmt.Println(segments2.ToString()) // [[20,1],[30,0]]
	segments2.Add(10, 40, -1)
	fmt.Println(segments2.ToString()) // [[10,-1],[20,0],[30,-1],[40,0]]
}
