package slicekit

import "math"

// Includes Includes[T comparable]
//
//	@Description: 切片中是否包含指定元素
//	@param s slice
//	@param target
//	@return bool
func Includes[T comparable](s []T, target T) bool {
	if len(s) == 0 {
		return false
	}

	for _, item := range s {
		if target == item {
			return true
		}
	}

	return false
}

// IncludesBy
//
//	@Description:	切片中是否包含 满足条件的 元素
//	@param s
//	@param predicate
//	@return bool
func IncludesBy[T interface{}](s []T, predicate func(item T) bool) bool {
	for _, item := range s {
		if predicate(item) {
			return true
		}
	}

	return false
}

// IncludesSubSlice IncludesSubSlice[T comparable{}]
//
//	@Description: 判断子切片 是否全部包含在 slice中
//	@param s slice
//	@param sub slice
//	@return bool
func IncludesSubSlice[T comparable](s []T, sub []T) bool {

	for _, item := range sub {
		if !Includes(s, item) {
			return false
		}
	}

	return true
}

// Each Each[T any]
//
//	@Description: 循环遍历
//	@param s
//	@param iteratee
func Each[T any](s []T, iteratee func(idx int, item T)) {
	for index, val := range s {
		iteratee(index, val)
	}
}

// Chunk Chunk[T any]
//
//	@Description: 按照指定size 切分数组
//	@param s
//	@param size
//	@return [][]T
func Chunk[T any](s []T, size int) [][]T {
	if size >= len(s) {
		r := make([][]T, 1, 1)
		r = append(r, s)
		return r
	}
	capSize := int(math.Ceil(float64(len(s) / size)))
	result := make([][]T, 0, capSize)

	for i := 0; i < capSize; i++ {
		startIdx := i * size
		endIndex := startIdx + size
		if endIndex >= len(s) {
			endIndex = len(s)
		}

		result = append(result, s[startIdx:endIndex])
	}

	return result
}

// Compact Compact[T comparable]
//
//	@Description:创建一个新数组，包含原数组中所有的非假值元素
//	@param s
//	@return []T
func Compact[T comparable](s []T) []T {
	var zero T

	result := make([]T, 0, len(s))
	for _, val := range s {
		if val != zero {
			result = append(result, val)
		}
	}

	return result
}

// Concat Concat[T any]
//
//	@Description: 创建一个新切片，将多个切片连接在一起
//	@param s
//	@param slices
//	@return []T
func Concat[T any](s []T, slices ...[]T) []T {
	result := append([]T{}, s...)

	for _, val := range slices {
		result = append(result, val...)
	}
	return result
}

// Difference Difference[T comparable]
//
//	@Description: 过滤掉指定元素 并返回一个新的切片
//	@param s
//	@param ignoreSlices
//	@return []T
func Difference[T comparable](s []T, ignoreSlices []T) []T {

	result := make([]T, 0, len(s))

	for _, val := range s {
		if !Includes(ignoreSlices, val) {
			result = append(result, val)
		}
	}

	return result
}

// DifferenceWith DifferenceWith[T comparable]
//
//	@Description: 自定义规则过滤掉指定元素 并返回一个新的切片
//	@param s
//	@param ignoreSlices
//	@param comparator
//	@return []T
func DifferenceWith[T comparable](s []T, ignoreSlices []T, comparator func(item1, item2 T) bool) []T {
	result := make([]T, 0, cap(s))

	checkHit := func(arr []T, it T, comparator func(item1, item2 T) bool) bool {
		isHit := false

		for _, v := range arr {
			if comparator(it, v) {
				isHit = true
				break
			}
		}

		return isHit
	}

	for _, val := range s {
		if !checkHit(ignoreSlices, val, comparator) {
			result = append(result, val)
		}
	}

	return result

}

// Equal Equal[T comparable]
//
//	@Description: 判断两个切片是否相等
//	@param s1
//	@param s2
//	@return bool
func Equal[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}

	for index, _ := range s1 {
		if s1[index] != s2[index] {
			return false
		}
	}

	return true
}

// EqualWith EqualWith[T comparable]
//
//	@Description: 按照指定规则 判断两个切片是否相等
//	@param s1
//	@param s2
//	@param comparator
//	@return bool
func EqualWith[T comparable](s1, s2 []T, comparator func(v1, v2 T) bool) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if !comparator(s1[i], s2[i]) {
			return false
		}
	}

	return true
}

// Filter Filter[T interface{}]
//
//	@Description: 筛选符合条件的切片
//	@param s
//	@param predicate
//	@return []T
func Filter[T interface{}](s []T, predicate func(idx int, item T) bool) []T {
	sliceList := make([]T, 0, cap(s))

	for index, val := range s {
		if predicate(index, val) {
			sliceList = append(sliceList, val)
		}
	}

	return sliceList
}

// FilterMap FilterMap[T any, D any]
//
//	@Description: 按照 规则 筛选并生成新的切片
//	@param s
//	@param iteratee
//	@return []D
func FilterMap[T any, D any](s []T, iteratee func(idx int, item T) (D, bool)) []D {
	result := make([]D, 0, cap(s))

	for index, val := range s {
		if res, ok := iteratee(index, val); ok {
			result = append(result, res)
		}
	}

	return result
}

// Map Map[T interface{}]
//
//	@Description: 按照规则 生成一个新数组
//	@param s
//	@param iteratee
//	@return []T
func Map[T interface{}](s []T, iteratee func(idx int, item T) T) []T {
	result := make([]T, len(s), cap(s))

	for index, val := range s {
		result[index] = iteratee(index, val)

	}
	return result
}
