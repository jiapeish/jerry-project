package main

import (
	"fmt"
	"strings"
)

const (
	Add = iota
	Set
)

type Segment struct {
	Left      int
	Right     int
	Intensity int
}

type IntensitySegments struct {
	segments []Segment
}

func NewIntensitySegments() *IntensitySegments {
	return &IntensitySegments{
		segments: make([]Segment, 0),
	}
}

func (is *IntensitySegments) Add(from, to, amount int) {
	is.update(from, to, amount, Add)
}

func (is *IntensitySegments) Set(from, to, amount int) {
	is.update(from, to, amount, Set)
}

func (is *IntensitySegments) update(from, to, amount int, operation int) {
	updated := make([]Segment, 0)
	n := len(is.segments)
	handledCount := 0

	if n == 0 {
		updated = append(updated, Segment{
			Left:      from,
			Right:     to,
			Intensity: amount,
		})
		is.segments = is.merge(updated)
		return
	}

	// 1. 当前扫描到的区间[left, right]在新变更的区间[from, to]左侧
	// left                      right         from                       to
	// ───────────────────────────────         ─────────────────────────────
	for i := 0; i < n && is.segments[i].Right <= from; i++ {
		updated = append(updated, is.segments[i])
		handledCount++
	}

	// 2. 当前扫描到的区间[left, right]与新变更的区间[from, to]有重叠
	//          from                       to
	//          ─────────────────────────────
	// left                                          right
	// ───────────────────────────────────────────────────
	for i := handledCount; i < n && is.segments[i].Left < to; i++ {
		current := is.segments[i]

		// 2.1 计算左边没重叠的部分
		//                from                       to
		//                ─────────────────────────────
		// left                      right
		// ───────────────────────────────
		if current.Left < from {
			updated = append(updated, Segment{
				Left:      current.Left,
				Right:     from,
				Intensity: current.Intensity,
			})
		}

		// 2.2 计算两个区间的重叠部分
		// from                       to
		// ─────────────────────────────
		//     left          right
		//     ───────────────────
		commonLeft := max(current.Left, from)
		commonRight := min(current.Right, to)
		intensity := 0
		if operation == Add {
			intensity = current.Intensity + amount
		} else if operation == Set {
			intensity = amount
		}
		updated = append(updated, Segment{
			Left:      commonLeft,
			Right:     commonRight,
			Intensity: intensity,
		})

		// 2.3 计算右边没重叠的部分
		// from                       to
		// ─────────────────────────────
		//           left                      right
		//           ───────────────────────────────
		if to < current.Right {
			updated = append(updated, Segment{
				Left:      to,
				Right:     current.Right,
				Intensity: current.Intensity,
			})
		}
		handledCount++
	}

	// 3. 当前扫描到的区间[left, right]在新变更的区间[from, to]右侧
	// from                       to     left                      right
	// ─────────────────────────────     ───────────────────────────────
	for i := handledCount; i < n; i++ {
		updated = append(updated, is.segments[i])
	}

	is.segments = is.merge(updated)
}

func (is *IntensitySegments) merge(segments []Segment) []Segment {
	if len(segments) <= 1 {
		return segments
	}

	merged := make([]Segment, 0)
	merged = append(merged, segments[0])

	for i := 1; i < len(segments); i++ {
		current := segments[i]
		lastIdx := len(merged) - 1
		last := merged[lastIdx]
		//          last
		// ─────────────────────────
		//                                 current
		//                          ───────────────────────
		if last.Right == current.Left && last.Intensity == current.Intensity {
			merged[lastIdx] = Segment{
				Left:      last.Left,
				Right:     current.Right,
				Intensity: last.Intensity,
			}
		} else {
			merged = append(merged, current)
		}
	}

	return merged
}

func (is *IntensitySegments) ToString() string {
	if len(is.segments) == 0 {
		return "[]"
	}
	parts := make([]string, len(is.segments))
	for i, seg := range is.segments {
		parts[i] = fmt.Sprintf("[%d,%d]", seg.Left, seg.Intensity)
	}
	return "[" + strings.Join(parts, ",") + "]"
}

func main() {
	segments := NewIntensitySegments()
	segments.Add(10, 30, 1)
	fmt.Println(segments.ToString()) // 输出 [[10,1],[30,0]]

	segments.Add(20, 40, 1)
	fmt.Println(segments.ToString()) // 输出 [[10,1],[20,2],[30,1],[40,0]]

	segments.Set(10, 40, -2)
	fmt.Println(segments.ToString()) // 输出 [[10,-1],[20,0],[30,-1],[40,0]]

	segments2 := NewIntensitySegments()
	segments2.Add(10, 30, 1)
	fmt.Println(segments2.ToString()) // 输出 [[10,1],[30,0]]

	segments2.Add(20, 40, 1)
	fmt.Println(segments2.ToString()) // 输出 [[10,1],[20,2],[30,1],[40,0]]

	segments2.Set(10, 40, -1)
	fmt.Println(segments2.ToString()) // 输出 [[20,1],[30,0]]

	segments2.Set(10, 40, -1)
	fmt.Println(segments2.ToString()) // 输出 [[10,-1],[20,0],[30,-1],[40,0]]
}
