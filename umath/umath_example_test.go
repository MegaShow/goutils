package umath

import "fmt"

func ExampleCeilFloat() {
	fmt.Println(CeilFloat(123.456, 0))
	fmt.Println(CeilFloat(123.456, 2))
	// Output:
	// 124
	// 123.46
}

func ExampleFloorFloat() {
	fmt.Println(FloorFloat(123.456, 0))
	fmt.Println(FloorFloat(123.456, 2))
	// Output:
	// 123
	// 123.45
}

func ExampleRoundFloat() {
	fmt.Println(RoundFloat(123.456, 0))
	fmt.Println(RoundFloat(123.456, 2))
	// Output:
	// 123
	// 123.46
}
