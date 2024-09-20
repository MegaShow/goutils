package uruntime

import "fmt"

func ExampleGetFuncFullName() {
	fmt.Println(GetFuncFullName(fmt.Println))
	fmt.Println(GetFuncFullName(GetFuncFullName))
	// Output:
	// fmt.Println
	// go.icytown.com/utils/uruntime.GetFuncFullName
}

func ExampleGetFuncName() {
	fmt.Println(GetFuncName(fmt.Println))
	fmt.Println(GetFuncName(GetFuncName))
	// Output:
	// Println
	// GetFuncName
}
