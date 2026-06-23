package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

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

func main() {
	envPath := flag.String("env", ".env", "path to .env file")
	examplePath := flag.String("example", ".env.example", "path to .env.example file")
	flag.Parse()
	example, err := parseEnvFile(*examplePath)

	if err != nil {
		fmt.Println("Error reading .env.example:", err)
		os.Exit(1)
	}

	local, err := parseEnvFile(*envPath)
	if err != nil {
		fmt.Println("Error reading .env:", err)
		os.Exit(1)
	}

	missingInLocal, missingInExample := diffKeys(example, local)

	for _, key := range missingInLocal {
		fmt.Println("x Missing in .env:           ", key)
	}
	for _, key := range missingInExample {
		fmt.Println("x Missing in .env.example:", key)
	}

	if len(missingInLocal) == 0 && len(missingInExample) == 0 {
		fmt.Println("All keys in sync")
		os.Exit(0)
	}

	os.Exit(1)
}