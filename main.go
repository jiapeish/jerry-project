package main

import (
	"fmt"
)

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
