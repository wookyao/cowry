package strkit

import (
	"math/rand"
	"time"
)

// makeRand 内部随机方法
func makeRand(ln int, targetStr string) string {
	cs := make([]byte, ln)

	for i := 0; i < ln; i++ {
		rand.Seed(time.Now().UnixNano())
		idx := rand.Intn(len(targetStr) - 1)
		cs[i] = targetStr[idx]
	}

	return string(cs)
}

// RandCus 随机自定义字符
func RandCus(ln int, targetStr string) string {
	return makeRand(ln, targetStr)
}

// RandNum 随机指定位数的数字 `0-9`
func RandNum(ln int) string {
	return makeRand(ln, "0123456789")
}

// RandChars 随机小写字符串
func RandChars(ln int) string {
	return makeRand(ln, "abcdefghijklmnopqrstuvwxyz")
}

// RandStr 随机小写字符串 `a-zA-z`
func RandStr(ln int) string {
	return makeRand(ln, "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz")
}

// RandCharsNum 随机字符串 `a-z0-9`
func RandCharsNum(ln int) string {
	return makeRand(ln, "abcdefghijklmnopqrstuvwxyz0123456789")
}

// RandCharsNum2 随机字符串 `a-zA-Z0-9`
func RandCharsNum2(ln int) string {
	return makeRand(ln, "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz0123456789")
}
