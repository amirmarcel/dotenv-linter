package main

import (
	"os"
	"strings"
)

func diffKeys(example, local map[string]bool) ([]string, []string) {
	var missingInLocal []string
	var missingInExample []string

	for key := range example {
		if !local[key] {
			missingInLocal = append(missingInLocal, key)
		}
	}

	for key := range local {
		if !example[key] {
			missingInExample = append(missingInExample, key)
		}
	}

	return missingInLocal, missingInExample
}

func parseEnvFile(path string) (map[string]bool, error) {
	data, err:= os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	keys:= make(map[string]bool)

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