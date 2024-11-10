package uslice

import "fmt"

func ExampleFind() {
	fmt.Println(Find([]int{1, 2, 3}, func(v int) bool { return v == 2 }))
	fmt.Println(Find([]int{1, 2, 3}, func(v int) bool { return v == 0 }))
	// Output:
	// 2 true
	// 0 false
}

func ExampleFindLast() {
	fmt.Println(FindLast([]int{1, 2, 3}, func(v int) bool { return v == 2 }))
	fmt.Println(FindLast([]int{1, 2, 3}, func(v int) bool { return v == 0 }))
	// Output:
	// 2 true
	// 0 false
}

func ExampleFilter() {
	a := []int{1, 2, 3}
	b := Filter(a, func(v int) bool {
		return v <= 2
	})
	fmt.Println(a)
	fmt.Println(b)
	// Output:
	// [1 2 3]
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

func ExampleUnique() {
	fmt.Println(Unique([]int{}))
	fmt.Println(Unique([]int{1, 1, 2}))
	fmt.Println(Unique([]int{1, 1, 2, 1, 3}))
	// Output:
	// []
	// [1 2]
	// [1 2 3]
}

func ExampleUniqueFunc() {
	a := []int{1, 1, 2, 1, 3}
	b := UniqueFunc(a, func(v int) int {
		return v % 2
	})
	fmt.Println(b)
	// Output:
	// [1 2]
}
