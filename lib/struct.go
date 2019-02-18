package lib

var baseType = map[string]bool{
	"int":         true,
	"int64":       true,
	"string":      true,
	"chan":        true,
	"map":         true,
	"float32":     true,
	"float64":     true,
	"bool":        true,
	"uint":        true,
	"int8":        true,
	"int16":       true,
	"int32":       true,
	"uint8":       true,
	"uint16":      true,
	"unit32":      true,
	"uint64":      true,
	"char":        true,
	"interface{}": true,
}

type EmptyStruct interface {
	Type()
}

type ImpStruct struct{}

func (ImpStruct) Type() {}

func getEmpty() EmptyStruct {
	var i ImpStruct
	return i
}
