package lib

import (
	"strings"
)

func coverRty(pack string, genMap map[string]GEN, rty string) string {
	if strings.Contains(rty, "[]") {
		crty := strings.Replace(rty, "[]", "", 1)
		if _, ok := genMap[crty]; ok {
			rty = "[]" + pack + "." + crty
		} else {
		}
	}

	if strings.Contains(rty, "[]*") {
		crty := strings.Replace(rty, "[]*", "", 1)
		if _, ok := genMap[crty]; ok {
			rty = "[]*" + pack + "." + crty
		}
	}

	if strings.Contains(rty, "*") {
		crty := strings.Replace(rty, "*", "", 1)
		if _, ok := genMap[crty]; ok {
			rty = "*" + pack + "." + crty
		}
	}

	return rty
}
