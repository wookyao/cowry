package cowrykit

import "fmt"

func ExampleToBool() {
	res1, _ := ToBool("1")
	res2, _ := ToBool("1 ")
	res3, _ := ToBool("")
	res4, _ := ToBool(false)
	res5, _ := ToBool([]string{"1"})

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	fmt.Println(res4)
	fmt.Println(res5)

	// Output:
	// true
	// true
	// false
	// false
	// false
}
