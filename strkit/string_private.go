package strkit

import (
	"fmt"
	"math"
	"regexp"
)

// calcPadStr
//
//	@Description: 计算需要填充字符串
//	@param diff 需要补充字符的个数
//	@param chars 指定的字符串
//	@return newChars
func calcPadStr(diff int, chars string) (newChars string) {
	// diff 小于等于 填充字符串的长度，只需将其截取即可
	if diff <= len(chars) {
		newChars = chars[:diff]
		return newChars
	}

	// count 填充字符需要 重复的次数
	count := math.Floor(float64(diff / len(chars)))

	// left 填充字符重复完， 剩余需要补充的个数
	left := diff % len(chars)

	newChars = repeatStr(chars, int(count)) + chars[:left]
	return
}

// repeatStr
//
//	@Description: 重复 N 次给定字符串
//	@param s 需要重复的字符串
//	@param count 重复次数
//	@return string
func repeatStr(s string, count int) string {
	if count <= 0 {
		return s
	}
	res := ""
	for i := 0; i < count; i++ {
		res += s
	}

	return res
}

// trim
//
//	@Description: 从string字符串中移除前面和后面的 空格 或 指定的字符
//	@param s 目标字符串
//	@param chars 指定字符
//	@param t [ 0 left&right, 1 left, 2 right, 3 all]
//	@return string
func trim(s, chars string, t int) string {
	if s == "" {
		return s
	}

	if chars == "" {
		chars = " "
	}

	regStr := ""
	if len(chars) > 0 {
		for _, val := range chars {
			regStr += fmt.Sprintf("\\%s|", string(val))
		}
	}

	var reg string
	regStrF := regStr[:len(regStr)-1]

	switch t {
	case 0: // left&right
		reg = fmt.Sprintf(`^(%s)*|(%s)*$`, regStrF, regStrF)
	case 1: // left
		reg = fmt.Sprintf(`^(%s)*`, regStrF)
	case 2: // right
		reg = fmt.Sprintf(`(%s)*$`, regStrF)
	default: // all
		reg = fmt.Sprintf(`(%s)*`, regStrF)
	}

	sampleRegexp := regexp.MustCompile(reg)

	result := sampleRegexp.ReplaceAllString(s, "")

	return result
}

// isLower
//
//	@Description: 判断字符是不是小写英文
//	@param r
//	@return bool
func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}
