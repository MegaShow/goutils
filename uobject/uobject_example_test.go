package uobject

import "fmt"

func ExampleDefault() {
	fmt.Println(Default(1, 100))
	fmt.Println(Default(0, 100))
	// Output:
	// 1
	// 100
}

func ExampleIndirect() {
	var v int = 1
	fmt.Println(Indirect(&v))
	fmt.Println(Indirect[int](nil))
	// Output:
	// 1
	// 0
}

func ExampleIndirectOr() {
	var v int = 1
	fmt.Println(IndirectOr(&v, 100))
	fmt.Println(IndirectOr[int](nil, 100))
	// Output:
	// 1
	// 100
}

func ExampleIsNotZero() {
	fmt.Println(IsNotZero(1))
	fmt.Println(IsNotZero(0))
	// Output:
	// true
	// false
}

func ExampleIsZero() {
	fmt.Println(IsZero(1))
	fmt.Println(IsZero(0))
	// Output:
	// false
	// true
}

func ExamplePtr() {
	fmt.Println(*Ptr(0))
	fmt.Println(*Ptr(1))
	// Output:
	// 0
	// 1
}
