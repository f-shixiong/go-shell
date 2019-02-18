package lib

import (
	"github.com/spf13/cast"
)

func BAdd(l interface{}, r interface{}) interface{} {
	switch l := l.(type) {
	case int:
		return l + cast.ToInt(r)
	case uint:
		return l + cast.ToUint(r)
	case int8:
		return l + cast.ToInt8(r)
	case int16:
		return l + cast.ToInt16(r)
	case int32:
		return l + cast.ToInt32(r)
	case int64:
		return l + cast.ToInt64(r)
	case uint8:
		return l + cast.ToUint8(r)
	case uint16:
		return l + cast.ToUint16(r)
	case uint32:
		return l + cast.ToUint32(r)
	case uint64:
		return l + cast.ToUint64(r)
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
	case uint:
		return l - cast.ToUint(r)
	case int8:
		return l - cast.ToInt8(r)
	case int16:
		return l - cast.ToInt16(r)
	case int32:
		return l - cast.ToInt32(r)
	case int64:
		return l - cast.ToInt64(r)
	case uint8:
		return l - cast.ToUint8(r)
	case uint16:
		return l - cast.ToUint16(r)
	case uint32:
		return l - cast.ToUint32(r)
	case uint64:
		return l - cast.ToUint64(r)
	case float32:
		return l - cast.ToFloat32(r)
	case float64:
		return l - cast.ToFloat64(r)
	default:
		Error("type= %#v", l)
	}
	return nil
}

func BMul(l interface{}, r interface{}) interface{} {
	switch l := l.(type) {
	case int:
		return l * cast.ToInt(r)
	case uint:
		return l * cast.ToUint(r)
	case int8:
		return l * cast.ToInt8(r)
	case int16:
		return l * cast.ToInt16(r)
	case int32:
		return l * cast.ToInt32(r)
	case int64:
		return l * cast.ToInt64(r)
	case uint8:
		return l * cast.ToUint8(r)
	case uint16:
		return l * cast.ToUint16(r)
	case uint32:
		return l * cast.ToUint32(r)
	case uint64:
		return l * cast.ToUint64(r)
	case float32:
		return l * cast.ToFloat32(r)
	case float64:
		return l * cast.ToFloat64(r)
	default:
		Error("type= %#v", l)
	}
	return nil
}

func BQuo(l interface{}, r interface{}) interface{} {
	switch l := l.(type) {
	case int:
		return l / cast.ToInt(r)
	case uint:
		return l / cast.ToUint(r)
	case int8:
		return l / cast.ToInt8(r)
	case int16:
		return l / cast.ToInt16(r)
	case int32:
		return l / cast.ToInt32(r)
	case int64:
		return l / cast.ToInt64(r)
	case uint8:
		return l / cast.ToUint8(r)
	case uint16:
		return l / cast.ToUint16(r)
	case uint32:
		return l / cast.ToUint32(r)
	case uint64:
		return l / cast.ToUint64(r)
	case float32:
		return l / cast.ToFloat32(r)
	case float64:
		return l / cast.ToFloat64(r)
	default:
		Error("type= %#v \n", l)
	}
	return nil
}

func BRem(l interface{}, r interface{}) interface{} {
	switch l := l.(type) {
	case int:
		return l % cast.ToInt(r)
	case uint:
		return l % cast.ToUint(r)
	case int8:
		return l % cast.ToInt8(r)
	case int16:
		return l % cast.ToInt16(r)
	case int32:
		return l % cast.ToInt32(r)
	case int64:
		return l % cast.ToInt64(r)
	case uint8:
		return l % cast.ToUint8(r)
	case uint16:
		return l % cast.ToUint16(r)
	case uint32:
		return l % cast.ToUint32(r)
	case uint64:
		return l % cast.ToUint64(r)
	default:
		Error("type= %#v \n", l)
	}
	return nil
}

func BAnd(l interface{}, r interface{}) interface{} {
	switch l := l.(type) {
	case int:
		return l & cast.ToInt(r)
	case uint:
		return l & cast.ToUint(r)
	case int8:
		return l & cast.ToInt8(r)
	case int16:
		return l & cast.ToInt16(r)
	case int32:
		return l & cast.ToInt32(r)
	case int64:
		return l & cast.ToInt64(r)
	case uint8:
		return l & cast.ToUint8(r)
	case uint16:
		return l & cast.ToUint16(r)
	case uint32:
		return l & cast.ToUint32(r)
	case uint64:
		return l & cast.ToUint64(r)
	default:
		Error("type = %#v", l)
	}
	return nil
}

func BOr(l interface{}, r interface{}) interface{} {
	switch l := l.(type) {
	case int:
		return l | cast.ToInt(r)
	case uint:
		return l | cast.ToUint(r)
	case int8:
		return l | cast.ToInt8(r)
	case int16:
		return l | cast.ToInt16(r)
	case int32:
		return l | cast.ToInt32(r)
	case int64:
		return l | cast.ToInt64(r)
	case uint8:
		return l | cast.ToUint8(r)
	case uint16:
		return l | cast.ToUint16(r)
	case uint32:
		return l | cast.ToUint32(r)
	case uint64:
		return l | cast.ToUint64(r)
	default:
		Error("type = %#v", l)
	}
	return false
}

func BXor(l interface{}, r interface{}) interface{} {
	switch l := l.(type) {
	case int:
		return l ^ cast.ToInt(r)
	case uint:
		return l ^ cast.ToUint(r)
	case int8:
		return l ^ cast.ToInt8(r)
	case int16:
		return l ^ cast.ToInt16(r)
	case int32:
		return l ^ cast.ToInt32(r)
	case int64:
		return l ^ cast.ToInt64(r)
	case uint8:
		return l ^ cast.ToUint8(r)
	case uint16:
		return l ^ cast.ToUint16(r)
	case uint32:
		return l ^ cast.ToUint32(r)
	case uint64:
		return l ^ cast.ToUint64(r)
	default:
		Error("type = %#v", l)
	}
	return nil
}

func BShl(l interface{}, r interface{}) interface{} {
	switch l := l.(type) {
	case uint:
		return l << cast.ToUint(r)
	case uint8:
		return l << cast.ToUint8(r)
	case uint16:
		return l << cast.ToUint16(r)
	case uint32:
		return l << cast.ToUint32(r)
	case uint64:
		return l << cast.ToUint64(r)
	default:
		Error("type = %#v", l)
	}
	return nil
}

func BShr(l interface{}, r interface{}) interface{} {
	switch l := l.(type) {
	case uint:
		return l >> cast.ToUint(r)
	case uint8:
		return l >> cast.ToUint8(r)
	case uint16:
		return l >> cast.ToUint16(r)
	case uint32:
		return l >> cast.ToUint32(r)
	case uint64:
		return l >> cast.ToUint64(r)
	default:
		Error("type = %#v", l)
	}
	return false
}

func BLss(l interface{}, r interface{}) bool {
	switch l := l.(type) {
	case int:
		return l < cast.ToInt(r)
	case uint:
		return l < cast.ToUint(r)
	case int8:
		return l < cast.ToInt8(r)
	case int16:
		return l < cast.ToInt16(r)
	case int32:
		return l < cast.ToInt32(r)
	case int64:
		return l < cast.ToInt64(r)
	case uint8:
		return l < cast.ToUint8(r)
	case uint16:
		return l < cast.ToUint16(r)
	case uint32:
		return l < cast.ToUint32(r)
	case uint64:
		return l < cast.ToUint64(r)
	case float32:
		return l < cast.ToFloat32(r)
	case float64:
		return l < cast.ToFloat64(r)
	default:
		Error("type = %#v", l)
	}
	return false
}

func BGtr(l interface{}, r interface{}) bool {
	switch l := l.(type) {
	case int:
		return l > cast.ToInt(r)
	case uint:
		return l > cast.ToUint(r)
	case int8:
		return l > cast.ToInt8(r)
	case int16:
		return l > cast.ToInt16(r)
	case int32:
		return l > cast.ToInt32(r)
	case int64:
		return l > cast.ToInt64(r)
	case uint8:
		return l > cast.ToUint8(r)
	case uint16:
		return l > cast.ToUint16(r)
	case uint32:
		return l > cast.ToUint32(r)
	case uint64:
		return l > cast.ToUint64(r)
	case float32:
		return l > cast.ToFloat32(r)
	case float64:
		return l > cast.ToFloat64(r)
	default:
		Error("type = %#v", l)
	}
	return false
}

func BLeq(l interface{}, r interface{}) bool {
	switch l := l.(type) {
	case int:
		return l <= cast.ToInt(r)
	case uint:
		return l <= cast.ToUint(r)
	case int8:
		return l <= cast.ToInt8(r)
	case int16:
		return l <= cast.ToInt16(r)
	case int32:
		return l <= cast.ToInt32(r)
	case int64:
		return l <= cast.ToInt64(r)
	case uint8:
		return l <= cast.ToUint8(r)
	case uint16:
		return l <= cast.ToUint16(r)
	case uint32:
		return l <= cast.ToUint32(r)
	case uint64:
		return l <= cast.ToUint64(r)
	case float32:
		return l <= cast.ToFloat32(r)
	case float64:
		return l <= cast.ToFloat64(r)
	default:
		Error("type = %#v", l)
	}
	return false
}

func BGeq(l interface{}, r interface{}) bool {
	switch l := l.(type) {
	case int:
		return l >= cast.ToInt(r)
	case uint:
		return l >= cast.ToUint(r)
	case int8:
		return l >= cast.ToInt8(r)
	case int16:
		return l >= cast.ToInt16(r)
	case int32:
		return l >= cast.ToInt32(r)
	case int64:
		return l >= cast.ToInt64(r)
	case uint8:
		return l >= cast.ToUint8(r)
	case uint16:
		return l >= cast.ToUint16(r)
	case uint32:
		return l >= cast.ToUint32(r)
	case uint64:
		return l >= cast.ToUint64(r)
	case float32:
		return l >= cast.ToFloat32(r)
	case float64:
		return l >= cast.ToFloat64(r)
	default:
		Error("type = %#v", l)
	}
	return false
}
