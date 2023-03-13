package syskit

import "fmt"

func ExampleGetOsBits() {
	b := GetOsBits()

	fmt.Println(b)
	// Output:
	// 64
}

func ExampleIsLinux() {
	isL := IsLinux()

	fmt.Println(isL)

	// Output:
	// false
}

func ExampleIsWindows() {
	isW := IsWindows()

	fmt.Println(isW)

	// Output:
	// false
}

func ExampleIsMac() {
	isM := IsMac()

	fmt.Println(isM)

	// Output:
	// true
}
