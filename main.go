package main

import (
	"encoding/json"
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
	quiet := flag.Bool("quiet", false, "suppress output, exit code only")
	jsonOut := flag.Bool("json", false, "output results as JSON")
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
	sort.Strings(missingInExample)

	if *jsonOut {
		result := map[string][]string{
			"missing_in_local":   missingInLocal,
			"missing_in_example": missingInExample,
		}
		data, _ := json.Marshal(result)
		fmt.Println(string(data))
	} else if !*quiet {
		for _, key := range missingInLocal {
			fmt.Println("x Missing in .env:           ", key)
		}
	
		
		for _, key := range missingInExample {
			fmt.Println("x Missing in .env.example:", key)
		}
	}

	if len(missingInLocal) == 0 && len(missingInExample) == 0 {
		if !*quiet && !*jsonOut {
			fmt.Println("All keys in sync")
		}
		os.Exit(0)
	}

	os.Exit(1)
}