package ucrypto

import "fmt"

func ExampleMD5Hex() {
	fmt.Println(MD5Hex("123456"))
	// Output:
	// e10adc3949ba59abbe56e057f20f883e
}

func ExampleSHA1Hex() {
	fmt.Println(SHA1Hex("123456"))
	// Output:
	// 7c4a8d09ca3762af61e59520943dc26494f8941b
}

func ExampleSHA256Hex() {
	fmt.Println(SHA256Hex("123456"))
	// Output:
	// 8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92
}

func ExampleSHA512Hex() {
	fmt.Println(SHA512Hex("123456"))
	// Output:
	// ba3253876aed6bc22d4a6ff53d8406c6ad864195ed144ab5c87621b6c233b548baeae6956df346ec8c17f5ea10f35ee3cbc514797ed7ddd3145464e2a0bab413
}
