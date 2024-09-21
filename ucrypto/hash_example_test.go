package ucrypto

import "fmt"

func ExampleMD5() {
	fmt.Println(MD5("123456"))
	// Output:
	// e10adc3949ba59abbe56e057f20f883e
}

func ExampleSHA512() {
	fmt.Println(SHA512("123456"))
	// Output:
	// ba3253876aed6bc22d4a6ff53d8406c6ad864195ed144ab5c87621b6c233b548baeae6956df346ec8c17f5ea10f35ee3cbc514797ed7ddd3145464e2a0bab413
}
