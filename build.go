package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	iconsDir, err := os.ReadDir("./assets")
	if err != nil {
		panic(err)
	}

	icons := make(map[string]string)

	for _, fileInfo := range iconsDir {
		if !fileInfo.IsDir() {
			data, err := os.ReadFile("./assets/" + fileInfo.Name())
			if err != nil {
				panic(err)
			}

			name := strings.TrimSuffix(strings.ToLower(fileInfo.Name()), ".svg")
			icons[name] = string(data)
		}
	}

	iconsJSON, err := json.Marshal(icons)
	if err != nil {
		panic(err)
	}

	indexFile, err := os.ReadFile("./api/icons.go")
	if err != nil {
		panic(err)
	}

	indexContent := string(indexFile)

	lines := strings.Split(indexContent, "\n")

	// Insert the iconsJSON into line 18
	lines[6] = fmt.Sprintf("\n\tvar iconsJSON string = `%s`\n", iconsJSON)

	// Join the lines back into a single string
	modifiedContent := strings.Join(lines, "\n")

	// Write the modified content back to the file
	err = os.WriteFile("./api/icons.go", []byte(modifiedContent), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Modified content saved to api/icons.go")
}
