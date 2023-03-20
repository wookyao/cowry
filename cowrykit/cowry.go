package cowrykit

import "fmt"

// ShortIF 模拟三目运算
func ShortIF[T interface{}](condition bool, ifThen, ifElse T) T {
	if condition {
		return ifThen
	} else {
		return ifElse
	}
}

// OrError 条件不匹配时，返回指定错误，否则返回nil
func OrError(condition bool, err error) error {
	if !condition {
		return err
	}
	return nil
}

// ErrOnFail OrError 的封装 => 条件不匹配时，返回指定错误，否则返回nil
func ErrOnFail(condition bool, err error) error {
	return OrError(condition, err)
}

// DataSize 获取数据大小
func DataSize(size float64) string {
	switch {
	case size < 1024:
		return fmt.Sprintf("%.0fB", size)
	case size < 1024*1024:
		return fmt.Sprintf("%.2fK", float64(size)/1024)
	case size < 1024*1024*1024:
		return fmt.Sprintf("%.2fM", float64(size)/1024/1024)
	default:
		return fmt.Sprintf("%.2fG", float64(size)/1024/1024/1024)

	}
}
