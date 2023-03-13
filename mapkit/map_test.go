package mapkit

import "fmt"

func ExampleKeys() {
	m := map[string]int{
		"go":     10,
		"java":   11,
		"python": 12,
	}

	keys := Keys(m)

	fmt.Println(keys)

	// Output:
	// [go java python]
}

func ExampleValues() {
	m := map[string]int{
		"go":     10,
		"java":   11,
		"python": 12,
	}

	values := Values(m)

	fmt.Println(values)

	// Output:
	// [10 11 12]
}

func ExampleMerge() {
	m1 := map[int]string{
		1: "a",
		2: "b",
	}
	m2 := map[int]string{
		1: "c",
		3: "d",
	}

	result := Merge(m1, m2)

	fmt.Println(result)

	// Output:
	// map[1:c 2:b 3:d]
}

func ExampleFilter() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	isEven := func(_ string, value int) bool {
		return value%2 == 0
	}

	result := Filter(m, isEven)

	fmt.Println(result)

	// Output:
	// map[b:2 d:4]
}

func ExampleFilterByKeys() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	result := FilterByKeys(m, []string{"a", "b"})

	fmt.Println(result)

	// Output:
	// map[a:1 b:2]
}

func ExampleFilterByValues() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	result := FilterByValues(m, []int{3, 4})

	fmt.Println(result)

	// Output:
	// map[c:3 d:4]
}

func ExampleOmit() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}
	isEven := func(_ string, value int) bool {
		return value%2 == 0
	}

	result := Omit(m, isEven)

	fmt.Println(result)

	// Output:
	// map[a:1 c:3 e:5]
}

func ExampleOmitByKeys() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	result := OmitByKeys(m, []string{"a", "b"})

	fmt.Println(result)

	// Output:
	// map[c:3 d:4 e:5]
}

func ExampleOmitByValues() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	result := OmitByValues(m, []int{4, 5})

	fmt.Println(result)

	// Output:
	// map[a:1 b:2 c:3]
}

func ExampleMinus() {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	m2 := map[string]int{
		"a": 11,
		"b": 22,
		"d": 33,
	}

	result := Minus(m1, m2)

	fmt.Println(result)

	// Output:
	// map[c:3]
}
