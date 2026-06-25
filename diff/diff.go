package diff

func DiffKeys(example, local map[string]bool) ([]string, []string) {
	missingInLocal := []string{}
	missingInExample := []string{}

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