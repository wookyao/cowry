package typekit

// IsString
//
//	@Description: 判断参数是否是string
//	@param v interface {}
//	@return bool
func IsString(v interface{}) bool {
	if v == nil {
		return false
	}
	switch v.(type) {
	case string:
		return true
	default:
		return false
	}
}
