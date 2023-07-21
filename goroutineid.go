// From: https://zhimin-wen.medium.com/logging-for-concurrent-go-programs-abe46ef14d58
// See also: https://github.com/golang/net/blob/master/http2/gotrack.go#L51

package goroutineid

import (
	"runtime"
	"strconv"
	"strings"
)

func Get() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
