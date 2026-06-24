package main

import (
	"flag"
	"fmt"
	"os"
)

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