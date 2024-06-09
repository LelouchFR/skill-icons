package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read directory contents
	iconsDir, err := os.ReadDir("./assets")
	if err != nil {
		panic(err)
	}

	// Initialize map for icons
	icons := make(map[string]string)

	// Iterate through directory contents
	for _, fileInfo := range iconsDir {
		if !fileInfo.IsDir() {
			// Read file contents
			data, err := os.ReadFile("./assets/" + fileInfo.Name())
			if err != nil {
				panic(err)
			}

			// Add file content to map
			name := strings.TrimSuffix(strings.ToLower(fileInfo.Name()), ".svg")
			icons[name] = string(data)
		}
	}

	// Convert icons to JSON
	iconsJSON, err := json.Marshal(icons)
	if err != nil {
		panic(err)
	}

	// Read content of api/index.go
	indexFile, err := os.ReadFile("./api/index.go")
	if err != nil {
		panic(err)
	}

	// Convert content to string
	indexContent := string(indexFile)

	// Split the content by newlines
	lines := strings.Split(indexContent, "\n")

	// Insert the iconsJSON into line 31
	lines[30] = fmt.Sprintf("var iconsJSON string = `%s`", iconsJSON)

	// Join the lines back into a single string
	modifiedContent := strings.Join(lines, "\n")

	// Write the modified content back to the file
	err = os.WriteFile("./api/index.go", []byte(modifiedContent), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Modified content saved to api/index.go")
}
