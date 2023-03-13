package validatorkit

import "fmt"

func ExampleCheck() {
	res1 := Check("123", letterRegexMatcher)

	fmt.Println(res1)
	// Output:
	// false
}

func ExampleIsHexColor() {
	res1 := IsHexColor("#000000")
	res2 := IsHexColor("#123456")
	res3 := IsHexColor("#1X2D3P")

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)

	// Output:
	// true
	// true
	// false
}

func ExampleIsAlpha() {
	res1 := IsAlpha("你好,jack")
	res2 := IsAlpha("你好,李明")
	res3 := IsAlpha("hello")

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)

	// Output:
	// false
	// false
	// true
}

func ExampleHasLetter() {
	res1 := HasLetter("你好,jack")
	res2 := HasLetter("你好,李明")
	res3 := HasLetter("hello")

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)

	// Output:
	// true
	// false
	// true
}

func ExampleIsIntStr() {
	res1 := IsIntStr("3")
	res2 := IsIntStr("-3")
	res3 := IsIntStr("3.")
	res4 := IsIntStr("3.0")
	res5 := IsIntStr("a")

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

func ExampleIsNumberStr() {
	res1 := IsNumberStr("3")
	res2 := IsNumberStr("-3")
	res3 := IsNumberStr("3.1415")
	res4 := IsNumberStr("3.0")
	res5 := IsNumberStr("4.")
	res6 := IsNumberStr("a")

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	fmt.Println(res4)
	fmt.Println(res5)
	fmt.Println(res6)

	// Output:
	// true
	// true
	// true
	// true
	// true
	// false
}

func ExampleIsIP() {
	res1 := IsIP("127.0.0.1")
	res2 := IsIPv4("::0:0:0:0:0:0:1")
	res3 := IsIPv6("127.0.0")
	res4 := IsIPv6("::0:0:0:0:")

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	fmt.Println(res4)

	// Output:
	// true
	// false
	// false
	// false
}

func ExampleIsPort() {
	res1 := IsPort("8080")
	res2 := IsPort("1525693")

	fmt.Println(res1)
	fmt.Println(res2)
	// Output:
	// true
	// false
}

func ExampleIsZeroValue() {
	type Stu struct {
	}

	res1 := IsZeroValue(1)
	res2 := IsZeroValue(false)
	res3 := IsZeroValue(3.25)
	res4 := IsZeroValue("")

	res5 := IsZeroValue(Stu{})

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	fmt.Println(res4)
	fmt.Println(res5)

	// Output:
	// false
	// true
	// false
	// true
	// true
}

func ExampleIsEmail() {
	result1 := IsEmail("abc@xyz.com")
	result2 := IsEmail("a.b@@com")

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleIsChineseMobile() {
	result1 := IsChineseMobile("15256936288")
	result2 := IsChineseMobile("12314589654")
	result3 := IsChineseMobile("05525621377")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// true
	// false
	// false
}

func ExampleIsChineseId() {
	result1 := IsChineseId("340322199101126035")
	result2 := IsChineseId("780322199101126035")
	result3 := IsChineseId("340322191001126035")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// true
	// true
	// true
}

func ExampleIsChinesePhone() {
	result1 := IsChinesePhone("010-32116675")
	result2 := IsChinesePhone("123-87562")

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleIsCreditCard() {
	result1 := IsCreditCard("4111111111111111")
	result2 := IsCreditCard("123456")

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleHasChinese() {
	result1 := HasChinese("abc")
	result2 := HasChinese("abC你好")
	result3 := HasChinese("你好")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// false
	// true
	// true
}

func ExampleIsBase64() {
	result1 := IsBase64("aGVsbG8=")
	result2 := IsBase64("123456")

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleIsJson() {
	res1 := IsJson("{\"age\": 15}")
	res2 := IsJson("{age: 15}")

	fmt.Println(res1)
	fmt.Println(res2)

	// Output:
	// true
	// false
}

func ExampleIsString() {
	res1 := IsString("123")
	res2 := IsString(123)

	fmt.Println(res1)
	fmt.Println(res2)

	// Output:
	// true
	// false
}

func ExampleIsAllUpper() {
	res1 := IsAllUpper("A1B2C3")
	res2 := IsAllUpper("ABC")
	res3 := IsAllUpper("abc")

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)

	// Output:
	// false
	// true
	// false
}

func ExampleIsAllLower() {
	res1 := IsAllLower("a1b2c3")
	res2 := IsAllLower("abc")
	res3 := IsAllLower("Abc")

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)

	// Output:
	// false
	// true
	// false
}

func ExampleHasUpperLetter() {
	res1 := HasUpperLetter("A1B2C3")
	res2 := HasUpperLetter("ABC")
	res3 := HasUpperLetter("abc")

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)

	// Output:
	// true
	// true
	// false
}

func ExampleHasLowerLetter() {
	res1 := HasLowerLetter("a1b2C3")
	res2 := HasLowerLetter("ABc")
	res3 := HasLowerLetter("ABC")

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)

	// Output:
	// true
	// true
	// false
}

func ExampleIsEmptyStr() {
	res1 := IsEmptyStr("")
	res2 := IsEmptyStr("1")

	fmt.Println(res1)
	fmt.Println(res2)

	// Output:
	// true
	// false
}

func ExampleIsStrongPassword() {
	res1 := IsStrongPassword("abcAbc123", 6)
	res2 := IsStrongPassword("Admin@125", 6)

	fmt.Println(res1)
	fmt.Println(res2)

	// Output:
	// false
	// true
}

func ExampleIsWeekPassword() {
	res1 := IsWeekPassword("abcAbc123")
	res2 := IsWeekPassword("Admin@125")

	fmt.Println(res1)
	fmt.Println(res2)

	// Output:
	// true
	// false
}
