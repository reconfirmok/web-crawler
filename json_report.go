package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func writeJSONReport(pages map[string]PageData, filename string) error {
	if len(pages) == 0 {
		fmt.Println("No data to write to JSON")
		return nil
	}

	pagesKeys := make([]string, 0, len(pages))
	for k := range pages {
		pagesKeys = append(pagesKeys, k)
	}
	sort.Strings(pagesKeys)

	pageData := make([]PageData, 0, len(pages))
	for _, k := range pagesKeys {
		pageData = append(pageData, pages[k])
	}

	data, err := json.MarshalIndent(pageData, "", "  ")
	if err != nil {
		return fmt.Errorf("Error marshaling json: %w", err)
	}

	err = os.WriteFile(filename, data, 0o644)
	if err != nil {
		return fmt.Errorf("Error writing JSON file: %w", err)
	}

	fmt.Printf("Report written to %s\n", filename)

	return nil
}
