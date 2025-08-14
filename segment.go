package main

func (is *IntensitySegments) Add(from, to, amount int) {
	if amount == 0 || from >= to {
		return
	}
	is.putDelta(from, amount)
	is.putDelta(to, -amount)
	
	is.update()
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

	is.update()
}

// ToString 打印
func (is *IntensitySegments) ToString() string {
	return is.cache
}
