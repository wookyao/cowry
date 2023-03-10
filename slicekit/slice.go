package slicekit

import (
	"github.com/wookyao/cowry/constraints"
	"math"
	"math/rand"
	"sort"
	"time"
)

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

	i, sLen := 0, len(s)
	for ; i < sLen; i++ {
		if target == s[i] {
			return true
		}
	}

	return false
}

// IncludesBy
//
//	@Description:	切片中是否包含 满足条件的 元素
//	@param s
//	@param comparator
//	@return bool
func IncludesBy[T interface{}](s []T, comparator func(item T) bool) bool {
	i, sLen := 0, len(s)
	for ; i < sLen; i++ {
		if comparator(s[i]) {
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
	i, sLen := 0, len(sub)
	for ; i < sLen; i++ {
		if !Includes(s, sub[i]) {
			return false
		}
	}

	return true
}

// Each Each[T any]
//
//	@Description: 循环遍历
//	@param s
//	@param comparator
func Each[T any](s []T, comparator func(idx int, item T)) {
	i, sLen := 0, len(s)
	for ; i < sLen; i++ {
		comparator(i, s[i])
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
		r := make([][]T, 0, 1)
		r = append(r, s)
		return r
	}
	capSize := int(math.Ceil(float64(len(s)) / float64(size)))
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

	i, sLen, result := 0, len(s), make([]T, 0, len(s))
	for ; i < sLen; i++ {
		if s[i] != zero {
			result = append(result, s[i])
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

// Every Every[T comparable]
//
//	@Description: 切片中所有元素都满足指定条件
//	@param s1
//	@param comparator
//	@return bool
func Every[T comparable](s []T, comparator func(idx int, item T) bool) bool {
	for idx, val := range s {
		if !comparator(idx, val) {
			return false
		}
	}

	return true
}

// None None[T any]
//
//	@Description: 切片中 没有满足指定条件的元素
//	@param s
//	@param comparator
//	@return bool
func None[T any](s []T, comparator func(idx int, item T) bool) bool {

	for i, v := range s {
		if comparator(i, v) {
			return false
		}
	}

	return true
}

// Some Some[T comparable]
//
//	@Description: 切片中有满足指定条件的元素
//	@param s
//	@param comparator
//	@return bool
func Some[T comparable](s []T, comparator func(idx int, item T) bool) bool {
	for idx, val := range s {
		if comparator(idx, val) {
			return true
		}
	}

	return false
}

// Filter Filter[T interface{}]
//
//	@Description: 筛选符合条件的切片
//	@param s
//	@param comparator
//	@return []T
func Filter[T interface{}](s []T, comparator func(idx int, item T) bool) []T {
	sliceList := make([]T, 0, cap(s))

	for index, val := range s {
		if comparator(index, val) {
			sliceList = append(sliceList, val)
		}
	}

	return sliceList
}

// FilterMap FilterMap[T any, D any]
//
//	@Description: 按照 规则 筛选并生成新的切片
//	@param s
//	@param comparator
//	@return []D
func FilterMap[T any, D any](s []T, comparator func(idx int, item T) (D, bool)) []D {
	result := make([]D, 0, cap(s))

	for index, val := range s {
		if res, ok := comparator(index, val); ok {
			result = append(result, res)
		}
	}

	return result
}

// Map Map[T interface{}]
//
//	@Description: 按照规则 生成一个新数组
//	@param s
//	@param comparator
//	@return []T
func Map[T interface{}](s []T, comparator func(idx int, item T) T) []T {
	result := make([]T, len(s), cap(s))

	for index, val := range s {
		result[index] = comparator(index, val)

	}
	return result
}

// Count Count[T comparable]
//
//	@Description: 计算指定元素在切片中出现的次数
//	@param s
//	@param target
//	@return int
func Count[T comparable](s []T, target T) int {
	count := 0

	for _, v := range s {
		if v == target {
			count++
		}
	}

	return count
}

// CountBy CountBy[T comparable]
//
//	@Description: 计算切片中 满足指定条件的元素个数
//	@param s
//	@param comparator
//	@return int
func CountBy[T comparable](s []T, comparator func(idx int, item T) bool) int {
	count := 0

	for i, v := range s {
		if comparator(i, v) {
			count++
		}
	}
	return count
}

// Find Find[T comparable]
//
//	@Description: 按照指定条件 从左至右筛选符合条件的地一个元素
//	@param s
//	@param comparator
//	@return *T
//	@return bool
func Find[T comparable](s []T, comparator func(idx int, item T) bool) (*T, bool) {
	index := -1

	for i, v := range s {
		if comparator(i, v) {
			index = i
			break
		}
	}

	if index == -1 {
		return nil, false
	}

	return &s[index], true
}

// FindLast FindLast[T comparable]
//
//	@Description: 按照指定条件 从右至左筛选符合条件的地一个元素
//	@param s
//	@param comparator
//	@return *T
//	@return bool
func FindLast[T comparable](s []T, comparator func(idx int, item T) bool) (*T, bool) {

	index := -1

	for i := len(s) - 1; i >= 0; i-- {
		if comparator(i, s[i]) {
			index = i
			break
		}
	}

	if index == -1 {
		return nil, false
	}
	return &s[index], true
}

// Fill Fill[T interface{}]
//
//	@Description: 填充指定长度，指定元素的切片
//	@param v
//	@param size
//	@return []T
func Fill[T interface{}](v T, size uint) []T {
	result := make([]T, 0, size)

	for i := 0; i < int(size); i++ {
		result = append(result, v)
	}

	return result
}

// Copy Copy[T any]
//
//	@Description: 浅拷贝
//	@param s
//	@return []T
func Copy[T any](s []T) []T {
	return sliceCopy(s)
}

// Sort - Sort[T constraints.Ordered]
//
// @Description: 排序
// @param s
// @param sortOrder
func Sort[T constraints.Ordered](s []T, sortOrder ...string) {
	// asc 升序 desc降序
	order := "desc"

	if len(sortOrder) > 0 && sortOrder[0] != "" {
		order = sortOrder[0]
	}

	sort.Slice(s, func(i, j int) bool {
		if order == "desc" {
			return s[i] > s[j]
		} else {
			return s[i] < s[j]
		}
	})
}

// SortBy SortBy[T any]
//
//	@Description: 根据指定规则排序
//	@param s
//	@param comparator
func SortBy[T any](s []T, comparator func(i, j int) bool) {
	sort.Slice(s, comparator)
}

// ToSlice ToSlice[T any]
//
//	@Description: 将输入参数转成slice
//	@param items
//	@return []T
func ToSlice[T any](items ...T) []T {
	var s = make([]T, 0, len(items))

	i, sLen := 0, len(items)

	for ; i < sLen; i++ {
		s = append(s, items[i])
	}

	return s
}

// Without[T comparable]
//
//	@Description: 剔除指定元素
//	@param s
//	@param items
//	@return []T
func Without[T comparable](s []T, items ...T) []T {
	if len(s) == 0 || len(items) == 0 {
		return s
	}

	i, sLen, list := 0, len(s), make([]T, 0, len(s))

	for ; i < sLen; i++ {
		if !Includes(items, s[i]) {
			list = append(list, s[i])
		}
	}

	return list
}

// IndexOf IndexOf[T comparable]
//
//	@Description: 查找目标元素首次出现的下标
//	@param s
//	@param item
//	@return int
func IndexOf[T comparable](s []T, item T) int {
	index := -1

	for i := 0; i < len(s); i++ {
		if s[i] == item {
			index = i
			break
		}
	}

	return index
}

// IndexOfLast IndexOfLast[T comparable]
//
//	@Description: 查找目标元素最后出现的下标
//	@param s
//	@param item
//	@return int
func IndexOfLast[T comparable](s []T, item T) int {
	index, i := -1, len(s)-1

	for ; i >= 0; i-- {
		if s[i] == item {
			index = i
			break
		}
	}

	return index
}

func Reverser[T any](s []T) {
	i, j := 0, len(s)-1

	for ; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Shuffle Shuffle[T any]
//
//	@Description: 随机打乱元素位置
//	@param s
func Shuffle[T any](s []T) {
	rand.Seed(time.Now().Unix()) // 取保每次的随机值与上一次不同

	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}

// Shift Shift[T any]
//
//	@Description: 不改变原切片 删除第一个元素, 返回新切片和删除的元素
//	@param s
//	@return []T
//	@return T
func Shift[T any](s []T) ([]T, T) {
	if len(s) == 0 {
		var t T
		return s, t
	}

	return s[1:], s[0]
}

// Pop Pop[T any]
//
//	@Description: 不改变原切片 方法从切片中删除最后一个元素，返回新切片和删除的元素
//	@param s
//	@return []T
//	@return T
func Pop[T any](s []T) ([]T, T) {
	if len(s) == 0 {
		var t T
		return s, t
	}

	return s[:len(s)-1], s[len(s)-1]
}

// Push Push[T any]
//
//	@Description: 不改变原切片 向切片尾部追加元素，返回新切片
//	@param s
//	@param items
//	@return []T
func Push[T any](s []T, items ...T) []T {
	return append(s, items...)
}

// Unshift Unshift[T any]
//
//	@Description: 不改变原切片 向切片头部追加元素，返回新切片
//	@param s
//	@param items
//	@return []T
func Unshift[T any](s []T, items ...T) []T {
	return append(items, s...)
}

// Insert Insert[T any]
//
//	@Description:不改变原切片 插入元素到指定下标，返回新切片
//	@param s
//	@param index
//	@param items
//	@return []T
func Insert[T any](s []T, index int, items ...T) []T {
	if len(items) == 0 {
		return s
	}

	if index >= len(s) {
		return Push(s, items...)
	} else if index <= 0 {
		return Unshift(s, items...)
	}

	left, right := sliceCopy(s[:index]), sliceCopy(s[index:])

	tmp := append(left, items...)
	tmp = append(tmp, right...)

	return tmp
}
