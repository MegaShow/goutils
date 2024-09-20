package uslice

import "fmt"

func ExampleDistinct() {
	fmt.Println(Distinct([]int{}))
	fmt.Println(Distinct([]int{1, 1, 2}))
	fmt.Println(Distinct([]int{1, 1, 2, 1, 3}))
	// Output:
	// []
	// [1 2]
	// [1 2 3]
}

func ExampleFind() {
	fmt.Println(Find([]int{1, 2, 3}, func(v int) bool { return v == 2 }))
	fmt.Println(Find([]int{1, 2, 3}, func(v int) bool { return v == 0 }))
	// Output:
	// 2 true
	// 0 false
}

func ExampleFilter() {
	a := []int{1, 2, 3}
	b := Filter(a, func(v int) bool {
		return v <= 2
	})
	fmt.Println(a) // a should not to be use.
	fmt.Println(b)
	// Output:
	// [1 2 0]
	// [1 2]
}

func ExampleGroupBy() {
	a := []int{1, 2, 3}
	m := GroupBy(a, func(v int) int {
		return v % 2
	})
	fmt.Println(m)
	// Output:
	// map[0:[2] 1:[1 3]]
}

func ExampleMap() {
	a := []int{1, 2, 3}
	b := Map(a, func(v int) int {
		return v * v
	})
	fmt.Println(b)
	// Output:
	// [1 4 9]
}

func ExampleOf() {
	fmt.Println(Of[int]())
	fmt.Println(Of(1))
	fmt.Println(Of(1, 2, 3))
	// Output:
	// []
	// [1]
	// [1 2 3]
}

func ExampleToMap() {
	a := []int{1, 2, 3}
	b := ToMap(a, func(v int) int {
		return v * v
	})
	fmt.Println(b)
	// Output:
	// map[1:1 4:2 9:3]
}
