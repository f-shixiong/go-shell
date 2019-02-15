package lib

import (
	"fmt"
)

var openDebug = false

func Debug(f string, args ...interface{}) {
	fix := "<Debug> "
	if openDebug {
		fmt.Printf(fix+f+"\n\n", args...)
	}
}

func Error(f string, args ...interface{}) {
	fix := "<Error> "
	fmt.Printf(fix+f+"\n\n", args...)
}

func Test(f string, args ...interface{}) {
	fix := "<Test> "
	fmt.Printf(fix+f+"\n\n", args...)
}
