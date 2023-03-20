package fmtkit

import "fmt"

func ExampleToChineseNumber() {
	res2 := ToChineseNumber(900195029009)

	fmt.Println(res2)

	// Output:
	// 九千零一亿九千五百零二万九千零九

}

func ExampleToAmountWords() {
	res1 := ToAmountWords(100025698000)

	fmt.Println(res1)

	// Output:
	// 壹仟亿贰仟伍佰陆拾玖万捌仟圆整
}
