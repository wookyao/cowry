package slicekit

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

// Filter Filter[T interface{}]
//
//	@Description: 筛选符合条件的切片
//	@param s
//	@param predicate
//	@return []T
func Filter[T interface{}](s []T, predicate func(idx int, item T) bool) []T {
	var sliceList []T

	for index, val := range s {
		if predicate(index, val) {
			sliceList = append(sliceList, val)
		}
	}

	return sliceList
}

// Map Map[T interface{}]
//
//	@Description: 按照规则 生成一个新数组
//	@param s
//	@param iteratee
//	@return []T
func Map[T interface{}](s []T, iteratee func(idx int, item T) T) []T {
	var sliceList []T

	for index, val := range s {
		result := iteratee(index, val)

		sliceList = append(sliceList, result)
	}
	return sliceList
}
