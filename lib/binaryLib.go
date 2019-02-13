package lib

import (
	"fmt"
)

func BAdd(l interface{}, r interface{}) interface{} {
	switch l := l.(type) {
	case int:
		return l + r.(int)
	default:
		fmt.Printf("anyway it happend,type= %#v \n", l)
	}
	return nil
}
