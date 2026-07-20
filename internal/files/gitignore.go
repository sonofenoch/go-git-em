package files

import (
	"os"
	"strings"
)

func ReadIgnoreList() []string {
	content, err := os.ReadFile("data.txt")
	if err != nil {
		// handle error
	}
	lines := strings.Split(string(content), "\n")
	return lines
}
