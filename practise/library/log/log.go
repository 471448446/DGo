package log

import (
	"fmt"
)

func Info(tag string, vars ...interface{}) {
	fmt.Println(tag+": ", vars)
}
