package lib

import (
	"testing"
)

func TestAutoImport(t *testing.T) {
	//TODO net/http cmd runtime strings 不支持

	//autoImport("net/http")
	autoImport("bufio")
}
