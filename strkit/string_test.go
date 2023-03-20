package strkit

import (
	"fmt"
	"testing"
)

func TestBefore(t *testing.T) {
	res1 := Before("Foo", "")
	res2 := Before("foo/bar", "/")
	res3 := Before("foo/bar/baz", "/baz")

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
}

func TestStartWith(t *testing.T) {
	res1 := StartWith("", "")
	res2 := StartWith("foo", "fo")
	res3 := StartWith("bar/foo", "bar/")
	res4 := StartWith("bar/foo", "baz")
	res5 := StartWith("barfoo", "barfoobaz")

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	fmt.Println(res4)
	fmt.Println(res5)
}

func TestEndWith(t *testing.T) {
	res1 := EndWith("", "")
	res2 := EndWith("foo", "o")
	res3 := EndWith("bar/foo", "foo")
	res4 := EndWith("bar/foo", "baz")
	res5 := EndWith("barfoo", "barfoobaz")
	res6 := EndWith("barfoo", "")

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	fmt.Println(res4)
	fmt.Println(res5)
	fmt.Println(res6)
}

func TestPad(t *testing.T) {
	pStr1 := Pad("abc", 8, "*_")
	pStr2 := Pad("abc", 5, "*_=$")
	pStr3 := Pad("abc", 10, "**_=$")

	res1 := PadStart("abc", 8, "**")
	res2 := PadStart("abc", 4, "**")
	res3 := PadStart("abc", 2, "*")

	fmt.Println(pStr1)
	fmt.Println(pStr2)
	fmt.Println(pStr3)

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
}

func TestTrim(t *testing.T) {
	var listStr map[string]string = map[string]string{
		"   abc ":      "",
		"-_-abc-_-":    "-_",
		"abc":          "",
		"-_-a-_-bc-_-": "-_",
	}

	for k, v := range listStr {
		fmt.Println("Trim:", Trim(k, v))
		fmt.Println("TrimLeft:", TrimLeft(k, v))
		fmt.Println("TrimRight:", TrimRight(k, v))
		fmt.Println("TrimAll:", TrimAll(k, v))
	}

}

func TestToLower(t *testing.T) {
	var strList = []string{"ABC", "Abc", "abc", "52AbC8Er"}

	for _, v := range strList {
		fmt.Printf("source:%s - result:%s \r\n", v, ToLower(v))
	}
}

func TestReverse(t *testing.T) {
	var strList = []string{"ABC", "Abc", "abc", "52AbC8Er"}
	for _, v := range strList {
		fmt.Printf("source:%s - result:%s \r\n", v, Reverse(v))
	}
}

func ExampleAt() {
	res1 := At("你好，世界", 0)
	res2 := At("hello word", 3)
	res3 := At("goland", 3)
	res4 := At("goland", -1)

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	fmt.Println(res4 == "")

	// Output:
	// 你
	// l
	// a
	// true
}
