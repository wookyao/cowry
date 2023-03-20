package fmtkit

import (
	"github.com/wookyao/cowry/cowrykit"
	"github.com/wookyao/cowry/slicekit"
	"regexp"
	"strconv"
)

// ToChineseNumber 数字转成中文
func ToChineseNumber(num int) string {
	return toChinesNum(num, false)
}

// ToAmountWords 数字转成大写金额
func ToAmountWords(num int) string {
	return toChinesNum(num, true) + "圆整"
}

func toChinesNum(num int, upper bool) string {
	if num <= 0 {
		return "零"
	}

	str := strconv.Itoa(num)
	var strList = make([]string, 0)

	for i := len(str); i >= 0; i = i - 4 {
		if i-4 > 0 {
			strList = append(strList, str[i-4:i])
		} else {
			strList = append(strList, str[:i])
		}
	}

	strList = slicekit.Filter(strList, func(idx int, item string) bool {
		return item != ""
	})

	result := ""
	units := []string{"", "万", "亿", "万亿"}
	hasM := false

	for idx, v := range strList {
		tv := tran(v, upper)
		if tv != "" {
			if idx == 2 {
				hasM = true
			}
			u := units[idx]
			if idx == 3 && hasM {
				u = "万"
			}
			result = tv + u + result
		}

	}

	return result
}

func tran(str string, upper bool) string {
	cNums := cowrykit.ShortIF(upper,
		[]string{"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖"},
		[]string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"})
	units := cowrykit.ShortIF(upper,
		[]string{"", "拾", "佰", "仟"},
		[]string{"", "十", "百", "千"})

	result := ""

	for i := len(str) - 1; i >= 0; i-- {
		v := string(str[i])
		idx, _ := strconv.Atoi(v)
		char := cNums[idx]
		uIdx := len(str) - i - 1
		u := units[uIdx]

		if result == "" {
			if char != cNums[0] {
				result = char + u + result
			}
		} else {
			if char == cNums[0] {
				if ok, _ := regexp.MatchString(cNums[0], result); !ok {
					result = char + result
				}

			} else {
				result = char + u + result
			}
		}

	}

	return result
}
