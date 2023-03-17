package strkit

import "fmt"

func ExampleRandCus() {
	result1 := RandCus(10, "56*/@#%^)polmhy")
	result2 := RandCharsNum2(2)
	fmt.Println(result1)
	fmt.Println(result2)
	// Output:
	//
}

func ExampleRandNum() {
	result1 := RandNum(1)
	result2 := RandNum(10)
	result3 := RandNum(21)
	fmt.Println(len(result1))
	fmt.Println(len(result2))
	fmt.Println(len(result3))

	// Output:
	//	1
	// 10
	// 21
}

func ExampleRandChars() {
	result1 := RandChars(100)

	fmt.Println(len(result1))

	// Output:
	// 100
}

func ExampleRandStr() {
	result1 := RandStr(12)

	fmt.Println(len(result1))

	// Output:
	// 12
}

func ExampleRandCharsNum() {
	result1 := RandCharsNum(15)
	fmt.Println(len(result1))

	// Output:
	// 15
}
