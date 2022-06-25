//go:build debug

package utils

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func DBG_PRINTF(format string, arg ...interface{}) {
	pc, fn, line, _ := runtime.Caller(1)
	function := runtime.FuncForPC(pc).Name()
	function = filepath.Base(function)

	fmt.Printf("<<File:%s  Line:%d  Function:%s>> ", fn, line, function)
	fmt.Printf(format, arg...)
}
