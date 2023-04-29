package puver

import (
	"strings"
)

func HashData(content string, os string, arch string) string {
	for _, line := range strings.Split(content, "\n") {
		if strings.Contains(line, os) && strings.Contains(line, arch) {
			return strings.Fields(line)[0]
		}
	}

	return ""
}
