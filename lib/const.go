package lib

import (
	"github.com/f-shixiong/go-shell/lib/internal/net"
	"github.com/f-shixiong/go-shell/lib/internal/strings"
	"plugin"
)

const (
	SELF = "self"
)

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
	"error":       true,
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

type ipack struct {
	tMap map[string]interface{}
	fMap map[string]interface{}
}

var internalPack = map[string]ipack{
	"strings": ipack{
		tMap: strings.StuMap,
		fMap: strings.FucMap,
	},
	"net": ipack{
		tMap: net.StuMap,
		fMap: net.FucMap,
	},
}

var importMap map[string]plugin.Symbol = make(map[string]plugin.Symbol, 0)
