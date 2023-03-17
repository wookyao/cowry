package cowrykit

import "fmt"

func ExampleHasChild() {
	result1 := HasChild("abc")
	result2 := HasChild([]string{"a", "b"})

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// false
	// true
}
