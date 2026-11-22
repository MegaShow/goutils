package umath

import "fmt"

func ExampleCeilFloat() {
	fmt.Println(CeilFloat(123.456, 0))
	fmt.Println(CeilFloat(123.456, 2))
	fmt.Println(CeilFloat(123.456, -1))
	// Output:
	// 124
	// 123.46
	// 130
}

func ExampleFloorFloat() {
	fmt.Println(FloorFloat(123.456, 0))
	fmt.Println(FloorFloat(123.456, 2))
	fmt.Println(FloorFloat(123.456, -1))
	// Output:
	// 123
	// 123.45
	// 120
}

func ExampleRoundFloat() {
	fmt.Println(RoundFloat(123.456, 0))
	fmt.Println(RoundFloat(123.456, 2))
	fmt.Println(RoundFloat(123.456, -1))
	// Output:
	// 123
	// 123.46
	// 120
}
