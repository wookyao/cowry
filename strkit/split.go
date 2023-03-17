package strkit

import "strings"

// Cut 字符串从左往右按照指定字符串分割 before, after, isFound
func Cut(s, sep string) (string, string, bool) {
	if idx := strings.Index(s, sep); idx >= 0 {
		return s[:idx], s[idx+len(sep):], true
	}

	return s, "", false
}

// LastCut  字符串从右往左按照指定字符串分割 before after isFound
func LastCut(s, sep string) (string, string, bool) {
	if idx := strings.LastIndex(s, sep); idx >= 0 {
		return s[:idx], s[idx+len(sep):], true
	}

	return s, "", false
}

// Split 分割字符串，并丢弃空字符串 返回string切片
func Split(s, sep string) []string {
	var ss = make([]string, 0)
	trimStr := TrimAll(s, " ")
	if trimStr == "" {
		return ss
	}

	for _, v := range strings.Split(s, sep) {
		if val := TrimAll(v, " "); val != "" {
			ss = append(ss, val)
		}
	}

	return ss
}
