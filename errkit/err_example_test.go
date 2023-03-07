package errkit

import (
	"fmt"
	"testing"
)

func TestNewErr(t *testing.T) {
	e := NewErr(NotFound, "用户不存在")
	fmt.Println(e.CodeToHttpStatus())
	fmt.Println(e.Error())

	// Output:
	// 404
	// 用户不存在
}
