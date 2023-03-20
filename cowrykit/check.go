package cowrykit

import (
	"bytes"
	"reflect"
)

// HasChild check element => [array slice map struct]
func HasChild(elem interface{}) bool {
	v := reflect.ValueOf(elem)
	switch v.Kind() {
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Struct:
		return true
	default:
		return false

	}

}

// IsArrayOrSlice 判断对象是否是 数组 或者 切片
func IsArrayOrSlice(elem interface{}) bool {
	v := reflect.ValueOf(elem)
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		return true
	default:
		return false
	}

}

// IsFunc 判断是不是函数
func IsFunc(elem interface{}) bool {
	v := reflect.ValueOf(elem)
	switch v.Kind() {
	case reflect.Func:
		return true
	default:
		return false
	}

}

// IsEqual 判断相比较的元素是否相等
func IsEqual(src, dst interface{}) bool {
	if src == nil || dst == nil {
		return src == dst
	}

	if IsFunc(src) || IsFunc(dst) {
		return false
	}

	bSrc, ok := src.([]byte)
	if !ok {
		return reflect.DeepEqual(src, dst)
	}

	bDst, ok := dst.([]byte)
	if !ok {
		return false
	}

	if bSrc == nil || bDst == nil {
		return bSrc == nil && bDst == nil
	}
	return bytes.Equal(bSrc, bDst)
}
