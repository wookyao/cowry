package cowrykit

import (
	"errors"
	"fmt"
)

func ExampleShortIF() {
	fmt.Println(ShortIF(5 > 4, "ok", "none"))
	fmt.Println(ShortIF(5 > 10, -1, ShortIF(5 < 3, 0, 1)))

	// Output:
	// ok
	// 1
}

func ExampleErrOnFail() {
	res1 := ErrOnFail(false, errors.New("fail"))

	fmt.Println(res1)

	// Output:
	// fail
}

func ExampleDataSize() {
	res1 := DataSize(1.2 * 1024 * 1024 * 1024)
	res2 := DataSize(100)

	fmt.Println(res1)
	fmt.Println(res2)

	// Output:
	// 1.20G
}
