package validatorkit

import (
	"strconv"
	"unicode"
)

// IsAlpha 校验字符串是否都由字母组成[a-zA-Z]
func IsAlpha(str string) bool {
	return alphaMatcher.MatchString(str)
}

// HasLetter 字符串中是否含有字母
func HasLetter(str string) bool {
	return letterRegexMatcher.MatchString(str)
}

// IsIntStr 验证字符串是否能转成int
func IsIntStr(str string) bool {
	return intStrMatcher.MatchString(str)
}

// IsFloatStr 验证字符串是否能转成float
func IsFloatStr(str string) bool {
	_, err := strconv.ParseFloat(str, 64)

	return err == nil
}

// IsNumberStr 验证字符串是否能转成数字
func IsNumberStr(str string) bool {
	return IsIntStr(str) || IsFloatStr(str)
}

// IsEmail 目标字符串是否符合email规则
func IsEmail(str string) bool {
	return emailMatcher.MatchString(str)
}

// IsChineseMobile 校验是否符合手机号码规则
func IsChineseMobile(str string) bool {
	return chineseMobileMatcher.MatchString(str)
}

// IsChineseId 简单版 验证字符串是否符合中国身份证号码
func IsChineseId(str string) bool {
	return chineseIdMatcher.MatchString(str)
}

// HasChinese 校验字符串中是否包含中文
func HasChinese(str string) bool {
	return chineseMatcher.MatchString(str)
}

// IsChinesePhone 校验固定电话
func IsChinesePhone(str string) bool {
	return chinesePhoneMatcher.MatchString(str)
}

// IsCreditCard 校验银行卡号
func IsCreditCard(str string) bool {
	return creditCardMatcher.MatchString(str)
}

// IsString 判断参数是否是string
func IsString(v interface{}) bool {
	if v == nil {
		return false
	}
	switch v.(type) {
	case string:
		return true
	default:
		return false
	}
}

// IsAllUpper 是否都是大写字母
func IsAllUpper(str string) bool {
	for _, v := range str {
		if !(v >= 'A' && v <= 'Z') {
			return false
		}
	}

	return true
}

// IsAllLower 是否都是小写字母
func IsAllLower(str string) bool {
	for _, v := range str {
		if !(v >= 'a' && v <= 'z') {
			return false
		}
	}

	return true
}

// HasUpperLetter 含有大写字母
func HasUpperLetter(str string) bool {
	for _, v := range str {
		if v >= 'A' && v <= 'Z' {
			return true
		}
	}

	return false
}

// HasLowerLetter 含有小写字母
func HasLowerLetter(str string) bool {
	for _, v := range str {
		if v >= 'a' && v <= 'z' {
			return true
		}
	}

	return false
}

// IsEmptyStr 是否是空字符串
func IsEmptyStr(str string) bool {
	return len(str) == 0
}

// IsStrongPassword 是否是强密码
func IsStrongPassword(str string, length int) bool {
	if len(str) < length {
		return false
	}

	hasNum, hasUpper, hasLower, hasSpec := false, false, false, false

	for _, v := range str {
		switch {
		case v >= 'A' && v <= 'Z':
			hasUpper = true
		case v >= 'a' && v <= 'z':
			hasLower = true
		case unicode.IsDigit(v):
			hasNum = true
		case unicode.IsSymbol(v) || unicode.IsPunct(v):
			hasSpec = true
		}
	}

	return hasSpec && hasUpper && hasLower && hasNum
}

// IsWeekPassword 弱密码判断
func IsWeekPassword(str string) bool {
	hasNum, hasLetter, hasSpec := false, false, false

	for _, v := range str {
		switch {
		case unicode.IsLetter(v):
			hasLetter = true
		case unicode.IsDigit(v):
			hasNum = true
		case unicode.IsSymbol(v) || unicode.IsPunct(v):
			hasSpec = true
		}
	}

	return (hasLetter || hasNum) && !hasSpec
}
