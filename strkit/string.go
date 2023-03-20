package strkit

import (
	"math"
	"strings"
)

// Before 返回源字符串中目标字符串首次出现之前的子字符串
func Before(s, chars string) string {
	if s == "" || chars == "" {
		return s
	}
	idx := strings.Index(s, chars)
	return s[:idx]
}

// BeforeLast 返回源字符串中目标字符串最后一次出现之前的子字符串
func BeforeLast(s, chars string) string {
	if s == "" || chars == "" {
		return s
	}
	idx := strings.LastIndex(s, chars)

	return s[:idx]
}

// After 返回源字符串中目标字符串首次出现之后的子字符串
func After(s, chars string) string {
	if s == "" || chars == "" {
		return s
	}
	idx := strings.Index(s, chars)
	return s[idx+len(chars):]

}

// AfterLast 返回源字符串中目标字符串最后一次出现之后的子字符串
func AfterLast(s, chars string) string {
	if s == "" || chars == "" {
		return s
	}
	idx := strings.LastIndex(s, chars)
	return s[idx+len(chars):]
}

// StartWith 检查字符串string是否以 target 开头。
func StartWith(s, chars string) bool {
	cLen := len(chars)
	if cLen > len(s) {
		return false
	}
	target := s[:cLen]

	return target == chars
}

// EndWith 检查字符串string是否以给定的target字符串结尾。
func EndWith(s, chars string) bool {
	sLen := len(s)
	cLen := len(chars)
	if cLen > sLen {
		return false
	}

	target := s[sLen-cLen:]

	return target == chars
}

// Pad 填充字符串
// 如果string字符串长度小于 length 则从左侧和右侧填充字符。 如果没法平均分配，则截断超出的长度。
func Pad(s string, length int, chars string) string {
	diff := length - len(s)
	if diff <= 0 {
		return s[:length]
	}

	newChars := calcPadStr(diff, chars)
	if len(newChars) <= len(chars) {
		return newChars + s
	}

	idx := int(math.Floor(float64(len(newChars) / 2)))
	if idx < len(chars) {
		idx = len(chars)
	}
	leftchars := newChars[:idx]
	rightchars := newChars[idx:]

	return leftchars + s + rightchars
}

// PadStart 填充左侧字符串
// 如果string字符串长度小于 length 则在左侧填充字符。 如果超出length长度则截断超出的部分。
func PadStart(s string, length int, chars string) string {
	diff := length - len(s)
	if diff <= 0 {
		return s[:length]
	}

	newChars := calcPadStr(diff, chars)
	return newChars + s

}

// PadEnd 填充右侧字符串
// 如果string字符串长度小于 length 则在右侧填充字符。 如果超出length长度则截断超出的部分。
func PadEnd(s string, length int, chars string) string {
	diff := length - len(s)
	if diff <= 0 {
		return s[:length]
	}

	newChars := calcPadStr(diff, chars)
	return s + newChars
}

// Repeat 重复 N 次给定字符串。
func Repeat(s string, count int) string {
	return repeatStr(s, count)
}

// Trim 从string字符串中移除前面和后面的 空格 或 指定的字符。
func Trim(s, chars string) string {
	return trim(s, chars, 0)
}

// TrimLeft 从string字符串中移除前面的 空格 或 指定的字符。
func TrimLeft(s, chars string) string {
	return trim(s, chars, 1)
}

// TrimRight 从string字符串中移除后面的 空格 或 指定的字符。
func TrimRight(s, chars string) string {
	return trim(s, chars, 2)
}

// TrimAll 从string字符串中移除所有的 空格 或 指定的字符。
func TrimAll(s, chars string) string {
	return trim(s, chars, 3)
}

// ToLower
//
//	@Description: 小写转大写
//	@param s
//	@return string
func ToLower(s string) string {
	if s == "" {
		return s
	}

	str := ""

	for _, v := range s {
		if isUpper(v) {
			str += string(v + 32)
		} else {
			str += string(v)
		}
	}

	return str
}

// LowerFirst
//
//	@Description: 转换字符串string的首字母为小写。
//	@param s
//	@return string
func LowerFirst(s string) string {
	if s == "" {
		return s
	}

	first := ToLower(s[0:1])

	return first + s[1:]
}

// ToUpper
//
//	@Description: 大写转小写
//	@param s
//	@return string
func ToUpper(s string) string {
	if s == "" {
		return s
	}

	str := ""
	for _, v := range s {
		if isUpper(v) {
			str += string(v - 32)
		} else {
			str += string(v)
		}
	}

	return str
}

// UpperFirst
//
//	@Description: 转换字符串string的首字母为大写。
//	@param s
//	@return string
func UpperFirst(s string) string {
	if s == "" {
		return s
	}

	first := ToUpper(s[0:1])

	return first + s[1:]
}

// Capitalize
//
//	@Description: 转换字符串string首字母为大写，剩下为小写。
//	@param s
//	@return string
func Capitalize(s string) string {
	if s == "" {
		return s
	}

	return ToUpper(s[:1]) + ToLower(s[1:])
}

// Reverse
//
//	@Description: 反转字符串
//	@param s 待反转的字符串
//	@return string
func Reverse(s string) string {
	r := []rune(s)
	i, j := 0, len(s)-1
	for ; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}

// Wrap
//
//	@Description: 指定字符 包裹string
//	@param s	目标字符串
//	@param chars 指定字符串
//	@return string
func Wrap(s, chars string) string {
	return Pad(s, len(s)+(len(chars)*2), chars)
}

// UnWrap
//
//	@Description: 清除包裹string的指定字符
//	@param s	目标字符串
//	@param chars 指定字符串
//	@return string
func UnWrap(s, chars string) string {
	return Trim(s, chars)
}

// At 获取指定下标的值
func At(str string, idx int) string {
	if len(str) < idx+1 {
		return ""
	}

	for index, v := range str {
		if idx == index {
			return string(v)
		}
	}

	return ""
}
