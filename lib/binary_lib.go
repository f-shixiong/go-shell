package lib

import (
	"github.com/spf13/cast"
)

func BAdd(l interface{}, r interface{}) interface{} {
	switch l := l.(type) {
	case int:
		return l + cast.ToInt(r)
	case int64:
		return l + cast.ToInt64(r)
	case float32:
		return l + cast.ToFloat32(r)
	case float64:
		return l + cast.ToFloat64(r)
	case string:
		return l + cast.ToString(r)

	default:
		Error("type= %#v \n", l)
	}
	return nil
}

func BSub(l interface{}, r interface{}) interface{} {
	switch l := l.(type) {
	case int:
		return l - cast.ToInt(r)
	case int64:
		return l - cast.ToInt64(r)
	case float32:
		return l - cast.ToFloat32(r)
	case float64:
		return l - cast.ToFloat64(r)
	default:
		Error("type= %#v \n", l)
	}
	return nil
}

func BMul(l interface{}, r interface{}) interface{} {
	switch l := l.(type) {
	case int:
		return l * cast.ToInt(r)
	case int64:
		return l * cast.ToInt64(r)
	case float32:
		return l * cast.ToFloat32(r)
	case float64:
		return l * cast.ToFloat64(r)
	default:
		Error("type= %#v \n", l)
	}
	return nil
}

func BQuo(l interface{}, r interface{}) interface{} {
	switch l := l.(type) {
	case int:
		return l / cast.ToInt(r)
	case int64:
		return l / cast.ToInt64(r)
	case float32:
		return l / cast.ToFloat32(r)
	case float64:
		return l / cast.ToFloat64(r)
	default:
		Error("type= %#v \n", l)
	}
	return nil
}

func BLss(l interface{}, r interface{}) bool {
	switch l := l.(type) {
	case int:
		return l < cast.ToInt(r)
	case int64:
		return l < cast.ToInt64(r)
	case float32:
		return l < cast.ToFloat32(r)
	default:
		Test("type %#v", l)
	}
	return false
}
