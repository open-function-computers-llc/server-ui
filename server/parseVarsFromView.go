package server

import "strings"

func parseVarsFromView(html string) (string, map[string]string) {
	output := html
	data := make(map[string]string)

	parts := strings.Split(html, "=====")
	if len(parts) == 0 {
		return output, data
	}

	if len(parts) != 3 {
		return output, data
	}

	output = strings.TrimSpace(parts[len(parts)-1])
	lines := strings.Split(parts[1], "\n")
	for _, line := range lines {
		vars := strings.Split(line, ":")
		if len(vars) != 2 {
			continue
		}
		data[strings.TrimSpace(vars[0])] = strings.TrimSpace(vars[1])
	}
	return output, data
}
