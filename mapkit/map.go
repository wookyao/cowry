package mapkit

import "github.com/wookyao/cowry/slicekit"

// Keys 获取map的key集合
func Keys[T comparable, V interface{}](m map[T]V) []T {
	keys := make([]T, 0)

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

// Values 获取map的value集合
func Values[T comparable, V interface{}](m map[T]V) []V {
	values := make([]V, 0)

	for _, v := range m {
		values = append(values, v)
	}

	return values
}

// Merge 合并多个map 后面的map覆盖之前的key
func Merge[T comparable, V interface{}](maps ...map[T]V) map[T]V {
	result := make(map[T]V, 0)

	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}

	return result
}

// Filter 筛选
func Filter[T comparable, V interface{}](m map[T]V, comparator func(k T, v V) bool) map[T]V {
	result := make(map[T]V, 0)

	for k, v := range m {
		if comparator(k, v) {
			result[k] = v
		}
	}

	return result
}

// FilterByKeys 根据指定keys 生成新的map
func FilterByKeys[T comparable, V interface{}](m map[T]V, keys []T) map[T]V {
	result := make(map[T]V, 0)

	for k, v := range m {
		if slicekit.Includes(keys, k) {
			result[k] = v
		}
	}

	return result
}

// FilterByValues 根据指定values 生成新的map
func FilterByValues[T comparable, V comparable](m map[T]V, values []V) map[T]V {
	result := make(map[T]V, 0)

	for k, v := range m {
		if slicekit.Includes(values, v) {
			result[k] = v
		}
	}

	return result
}

// Omit 丢去符合条件的元素
func Omit[T comparable, V interface{}](m map[T]V, comparator func(k T, v V) bool) map[T]V {
	result := make(map[T]V, 0)

	for k, v := range m {
		if !comparator(k, v) {
			result[k] = v
		}
	}

	return result
}

// OmitByKeys 剔除指定keys 生成新的map
func OmitByKeys[T comparable, V interface{}](m map[T]V, keys []T) map[T]V {
	result := make(map[T]V, 0)

	for k, v := range m {
		if !slicekit.Includes(keys, k) {
			result[k] = v
		}
	}

	return result
}

// OmitByValues 剔除指定values 生成新的map
func OmitByValues[T comparable, V comparable](m map[T]V, values []V) map[T]V {
	result := make(map[T]V, 0)

	for k, v := range m {
		if !slicekit.Includes(values, v) {
			result[k] = v
		}
	}

	return result
}

// Minus 以m1为基础 丢弃所有在 m2 存在的key 返回一个新的map
func Minus[T comparable, V interface{}](m1, m2 map[T]V) map[T]V {
	result := make(map[T]V, 0)

	for k, v := range m1 {
		if _, ok := m2[k]; !ok {
			result[k] = v
		}
	}

	return result
}
