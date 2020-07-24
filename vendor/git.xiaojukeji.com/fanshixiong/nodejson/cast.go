package nodejson

import (
	"encoding/json"
	"errors"
	"reflect"
)

func (n Node) Int() int {
	v, _ := n.IntE()
	return v
}

func (n Node) Int64() int64 {
	v, _ := n.Int64E()
	return v
}

func (n Node) String() string {
	v, _ := n.StringE()
	return v
}
func (n Node) Bool() bool {
	v, _ := n.BoolE()
	return v
}

func (n Node) Interface() interface{} {
	return n.data
}

func (n *Node) IntE() (int, error) {
	switch n.data.(type) {
	case json.Number:
		i, err := n.data.(json.Number).Int64()
		return int(i), err
	case float32, float64:
		return int(reflect.ValueOf(n.data).Float()), nil
	case int, int8, int16, int32, int64:
		return int(reflect.ValueOf(n.data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return int(reflect.ValueOf(n.data).Uint()), nil
	}
	return 0, errors.New("invalid value type")
}

func (n *Node) Float64E() (float64, error) {
	switch n.data.(type) {
	case json.Number:
		return n.data.(json.Number).Float64()
	case float32, float64:
		return reflect.ValueOf(n.data).Float(), nil
	case int, int8, int16, int32, int64:
		return float64(reflect.ValueOf(n.data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return float64(reflect.ValueOf(n.data).Uint()), nil
	}
	return 0, errors.New("invalid value type")
}

func (n *Node) Int64E() (int64, error) {
	switch n.data.(type) {
	case json.Number:
		return n.data.(json.Number).Int64()
	case float32, float64:
		return int64(reflect.ValueOf(n.data).Float()), nil
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(n.data).Int(), nil
	case uint, uint8, uint16, uint32, uint64:
		return int64(reflect.ValueOf(n.data).Uint()), nil
	}
	return 0, errors.New("invalid value type")
}

func (n *Node) StringE() (string, error) {
	if s, ok := (n.data).(string); ok {
		return s, nil
	}
	return "", errors.New("type assertion to string failed")

}

func (n *Node) BoolE() (bool, error) {
	if s, ok := (n.data).(bool); ok {
		return s, nil
	}
	return false, errors.New("type assertion to bool failed")
}
