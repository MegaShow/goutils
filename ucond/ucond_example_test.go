package ucond

import "fmt"

func ExampleIf() {
	fmt.Println(If(true, 1, 0))
	fmt.Println(If(false, 1, 0))
	// Output:
	// 1
	// 0
}

func ExampleIfFunc() {
	fmt.Println(IfFunc(true, func() int { return 1 }, func() int { return 0 }))
	fmt.Println(IfFunc(false, func() int { return 1 }, func() int { return 0 }))
	// Output:
	// 1
	// 0
}
