package lib

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"
)

var openDebug = false

func Debug(f string, args ...interface{}) {
	fix := "<Debug> "
	if openDebug {
		fmt.Printf(fix+f+"\n\n", args...)
	}
}

func Error(f string, args ...interface{}) {
	fix := "<Error> ["
	stack := getStack()
	fix += stack + "] "
	fmt.Printf(fix+f+"\n\n", args...)
	os.Exit(0)
}

func Test(f string, args ...interface{}) {
	fix := "<Test> "
	fmt.Printf(fix+f+"\n\n", args...)
}

func getStack() string {
	str := string(debug.Stack())
	arr := strings.Split(str, "\n")
	if len(arr) > 9 {
		return strings.Replace(arr[8], "\t", "", 20)
	}
	return ""
}
