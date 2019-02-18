package lib

import (
	"reflect"
)

func getResult(index int, ret interface{}) interface{} {
	Debug("index = %#v,ret = %#v", index, ret)
	if ret == nil {
		return nil
	}
	if index == -1 || reflect.TypeOf(ret).Kind() != reflect.Slice {
		return ret
	}
	rs := ret.([]interface{})
	if len(rs) <= index {
		return ret
	}
	return ret.([]interface{})[index]
}
