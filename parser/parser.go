package parser

import (
	"os"
	"strings"
)

func ParseEnvFile(path string) (map[string]bool, error){
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	keys := make(map[string]bool)

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) >= 1 {
			keys[strings.TrimSpace(parts[0])] = true
		}
	}

	return keys, nil
}