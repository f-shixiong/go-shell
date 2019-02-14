package lib

import (
	"fmt"
)

var openDebug = false

func Debug(f string, args ...interface{}) {
	if openDebug {
		fmt.Printf(f+"\n\n", args...)
	}
}

func Error(f string, args ...interface{}) {
	fmt.Printf(f+"\n\n", args...)
}
