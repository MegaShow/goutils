package uversion

import "fmt"

func ExampleNewSemVer() {
	v1, _ := NewSemVer(1, 2, 3, "", "")
	v2, _ := NewSemVer(1, 2, 3, "alpha", "fd")
	fmt.Println(v1)
	fmt.Println(v2)
	// Output:
	// 1.2.3
	// 1.2.3-alpha+fd
}

func ExampleMustNewSemVer() {
	fmt.Println(MustNewSemVer(1, 2, 3, "", ""))
	fmt.Println(MustNewSemVer(1, 2, 3, "alpha", "fd"))
	// Output:
	// 1.2.3
	// 1.2.3-alpha+fd
}

func ExampleParseSemVer() {
	v1, _ := ParseSemVer("1.2.3")
	v2, _ := ParseSemVer("1.2.3-alpha+fd")
	fmt.Println(v1)
	fmt.Println(v2)
	// Output:
	// 1.2.3
	// 1.2.3-alpha+fd
}

func ExampleMustParseSemVer() {
	fmt.Println(MustParseSemVer("1.2.3"))
	fmt.Println(MustParseSemVer("1.2.3-alpha+fd"))
	// Output:
	// 1.2.3
	// 1.2.3-alpha+fd
}

func ExampleSemVer_Major() {
	v := MustParseSemVer("1.2.3-alpha+fd")
	fmt.Println(v.Major())
	// Output:
	// 1
}

func ExampleSemVer_Minor() {
	v := MustParseSemVer("1.2.3-alpha+fd")
	fmt.Println(v.Minor())
	// Output:
	// 2
}

func ExampleSemVer_Patch() {
	v := MustParseSemVer("1.2.3-alpha+fd")
	fmt.Println(v.Patch())
	// Output:
	// 3
}

func ExampleSemVer_PreRelease() {
	v := MustParseSemVer("1.2.3-alpha+fd")
	fmt.Println(v.PreRelease())
	// Output:
	// alpha
}

func ExampleSemVer_Build() {
	v := MustParseSemVer("1.2.3-alpha+fd")
	fmt.Println(v.Build())
	// Output:
	// fd
}

func ExampleSemVer_String() {
	v := MustParseSemVer("1.2.3-alpha+fd")
	fmt.Println(v.String())
	// Output:
	// 1.2.3-alpha+fd
}

func ExampleSemVer_IncrMajor() {
	v := MustParseSemVer("1.2.3-alpha+fd")
	v.IncrMajor()
	fmt.Println(v)
	// Output:
	// 2.0.0
}

func ExampleSemVer_IncrMinor() {
	v := MustParseSemVer("1.2.3-alpha+fd")
	v.IncrMinor()
	fmt.Println(v)
	// Output:
	// 1.3.0
}

func ExampleSemVer_IncrPatch() {
	v := MustParseSemVer("1.2.3-alpha+fd")
	v.IncrPatch()
	fmt.Println(v)
	// Output:
	// 1.2.4
}

func ExampleSemVer_SetMajor() {
	v := MustParseSemVer("1.2.3-alpha+fd")
	v.SetMajor(5)
	fmt.Println(v)
	// Output:
	// 5.2.3-alpha+fd
}

func ExampleSemVer_SetMinor() {
	v := MustParseSemVer("1.2.3-alpha+fd")
	v.SetMinor(5)
	fmt.Println(v)
	// Output:
	// 1.5.3-alpha+fd
}

func ExampleSemVer_SetPatch() {
	v := MustParseSemVer("1.2.3-alpha+fd")
	v.SetPatch(5)
	fmt.Println(v)
	// Output:
	// 1.2.5-alpha+fd
}

func ExampleSemVer_SetPreRelease() {
	v := MustParseSemVer("1.2.3-alpha+fd")
	_ = v.SetPreRelease("beta")
	fmt.Println(v)
	// Output:
	// 1.2.3-beta+fd
}
func ExampleSemVer_SetBuild() {
	v := MustParseSemVer("1.2.3-alpha+fd")
	_ = v.SetBuild("ff")
	fmt.Println(v)
	// Output:
	// 1.2.3-alpha+ff
}

func ExampleSemVer_Compare() {
	v1 := MustParseSemVer("1.0.0-alpha")
	v2 := MustParseSemVer("1.0.0")
	fmt.Println(v1.Compare(v2))
	// Output:
	// -1
}
