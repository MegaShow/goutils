package usync

import "fmt"

func ExampleSingleflight() {
	var sf Singleflight[int, string]
	value, err := sf.Do(1, func() (string, error) {
		return "1", nil
	})
	fmt.Println(value, err)
	// Output:
	// 1 <nil>
}
