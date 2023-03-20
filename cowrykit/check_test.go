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

func ExampleIsEqual() {
	result1 := IsEqual("123", "123")
	result2 := IsEqual([]string{"123"}, "123")
	result3 := IsEqual([]string{"123"}, []string{"123"})
	result4 := IsEqual([]string{"123"}, []string{"12"})
	result5 := IsEqual(struct {
		Name string
	}{"123"}, struct {
		Name string
	}{"123"})
	result6 := IsEqual(struct {
		Name string
	}{"123"}, struct {
		Name string
	}{"12"})

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)

	// Output:
	// true
	// false
	// true
	// false
	// true
	// false
}

func fn1() {

}

func ExampleIsFunc() {
	var fn = func() {}

	res1 := IsFunc(fn)
	res2 := IsFunc(fn1)
	res3 := IsFunc("fn1")

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)

	// Output:
	// true
	// true
	// false
}

func ExampleIsArrayOrSlice() {
	res1 := IsArrayOrSlice([1]string{"1"})
	res2 := IsArrayOrSlice([]int{1, 2})
	res3 := IsArrayOrSlice(fn1)

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)

	// Output:
	// true
	// true
	// false
}
