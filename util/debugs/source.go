package debugs

import (
	"fmt"
	"runtime"
	"strings"
)

// SourceCodeLoc to get source code location
// example: SourceCodeLoc(1) to get current line number
func SourceCodeLoc(callDepth int) string {
	_, file, line, ok := runtime.Caller(callDepth)
	if !ok {
		return ""
	}
	file = strings.ReplaceAll(file, "\\", "/")
	paths := strings.Split(file, "/")
	if len(paths) > 3 {
		file = strings.Join(paths[len(paths)-3:], "/")
	}
	return fmt.Sprintf("%s:%d", file, line)
}

// WarpError warp line number info to error
func WarpError(err error, infos ...string) error {
	return fmt.Errorf("[%s] %+v, err=%+v", SourceCodeLoc(2), infos, err)
}
