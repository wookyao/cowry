package cowrykit

import "fmt"

func ExampleShortIF() {
	fmt.Println(ShortIF(5 > 4, "ok", "none"))
	fmt.Println(ShortIF(5 > 10, -1, ShortIF(5 < 3, 0, 1)))

	// Output:
	// ok
	// 1
}
