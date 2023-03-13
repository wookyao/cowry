package validatorkit

import (
	"encoding/json"
	"net"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// Check 自定义验证正则
func Check(str string, reg *regexp.Regexp) bool {
	return reg.MatchString(str)
}

// IsHexColor 十六进制颜色
func IsHexColor(str string) bool {
	return hexColorMatcher.MatchString(str)
}

// IsBase64 是否符合base64
func IsBase64(str string) bool {
	return base64Matcher.MatchString(str)
}

// IsJson 判断字符串是否可以转换成json
func IsJson(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// IsIP 验证字符串是否是ip地址
func IsIP(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil
}

// IsIPv4 是否满足IPv4规则
func IsIPv4(str string) bool {
	ip := net.ParseIP(str)
	if ip == nil {
		return false
	}

	return strings.Contains(str, ".")
}

// IsIPv6 是否满足IPv6规则
func IsIPv6(str string) bool {
	ip := net.ParseIP(str)
	if ip == nil {
		return false
	}

	return strings.Contains(str, ":")
}

// IsPort 字符串是否可以满足网络端口
func IsPort(str string) bool {
	port, err := strconv.Atoi(str)
	if err != nil {
		return false
	}

	return port > 0 && port < 65536
}

// IsZeroValue 是否是零值
func IsZeroValue(val interface{}) bool {
	if val == nil {
		return true
	}

	rv := reflect.ValueOf(val)

	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if !rv.IsValid() {
		return true
	}

	switch rv.Kind() {
	case reflect.String:
		return rv.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() == 0
	case reflect.Float64, reflect.Float32:
		return rv.Float() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return rv.Uint() == 0
	case reflect.Bool:
		return !rv.Bool()
	case reflect.Chan, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Func:
		return rv.IsNil()

	}
	return reflect.DeepEqual(rv.Interface(), reflect.Zero(rv.Type()).Interface())
}
