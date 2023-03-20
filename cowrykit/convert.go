package cowrykit

import (
	"errors"
	"fmt"
	"github.com/wookyao/cowry/strkit"
)

// ToBool 将目标元素转成bool
func ToBool(v interface{}) (bool, error) {
	if bv, ok := v.(bool); ok {
		return bv, nil
	}

	if str, ok := v.(string); ok {
		return Str2Bool(str)
	}

	return false, errors.New("convert value type error")

}

// Str2Bool 字符转成bool
func Str2Bool(v string) (bool, error) {
	switch strkit.ToLower(v) {
	case "1", "true":
		return true, nil
	case "0", "false":
		return false, nil

	}

	return false, fmt.Errorf("'%s' cannot convert to bool", v)
}
