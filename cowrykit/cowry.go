package cowrykit

// ShortIF 模拟三目运算
func ShortIF(condition bool, ifThen, ifElse interface{}) interface{} {
	if condition {
		return ifThen
	} else {
		return ifElse
	}
}
