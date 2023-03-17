package strkit

import "fmt"

func ExampleCut() {
	b1, a1, _ := Cut("123.456", ".")
	b2, a2, _ := Cut("123.456", "5")
	b3, a3, _ := Cut("123.456", "3")
	b4, a4, _ := Cut("123.456", "6")

	fmt.Println(b1, a1)
	fmt.Println(b2, a2)
	fmt.Println(b3, a3)
	fmt.Println(b4, a4)

	// Output:
	// 123 456
	// 123.4 6
	// 12 .456
	// 12.45

}

func ExampleLastCut() {
	b1, a1, _ := LastCut("12.34.56", ".")
	b2, a2, _ := LastCut("12.34.56", "4")
	b3, a3, _ := LastCut("12.34.56", "89")

	fmt.Println(b1, a1)
	fmt.Println(b2, a2)
	fmt.Println(b3, a3)

	// Output:
	//12.34 56
	//12.3 .56
	//12.34.56
}

func ExampleSplit() {
	res1 := Split("a,b,c, ,d", ",")
	res2 := Split("a,b,c, ,d", ".")

	fmt.Println(res1)
	fmt.Println(res2)

	// Output:
	// [a b c d]
}
