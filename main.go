package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/amirmarcel/dotenv-linter/diff"
	"github.com/amirmarcel/dotenv-linter/parser"
)

func main() {
	envPath := flag.String("env", ".env", "path to .env file")
	examplePath := flag.String("example", ".env.example", "path to .env.example file")
	flag.Parse()

	example, err := parser.ParseEnvFile(*examplePath)
	if err != nil {
		fmt.Println("Error reading .env.example:", err)
		os.Exit(1)
	}

	local, err := parser.ParseEnvFile(*envPath)
	if err != nil {
		fmt.Println("Error reading .env:", err)
		os.Exit(1)
	}

	missingInLocal, missingInExample := diff.DiffKeys(example, local)

	sort.Strings(missingInLocal)
	for _, key := range missingInLocal {
		fmt.Println("x Missing in .env:           ", key)
	}

	sort.Strings(missingInExample)
	for _, key := range missingInExample {
		fmt.Println("x Missing in .env.example:", key)
	}

	if len(missingInLocal) == 0 && len(missingInExample) == 0 {
		fmt.Println("All keys in sync")
		os.Exit(0)
	}

	os.Exit(1)
}