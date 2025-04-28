package library

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// TemplateData struct to hold the structure of your JSON data.
type TemplateData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Initiate Json Function
func InitiateJsonFunction(flags Flag) {
	// Find Matching Name and Description.
	if *flags.SearchTemplate && *flags.Keyword != "" {
		matchingNames, err := findMatchingTemplates(*flags.Path, *flags.Keyword)
		if err != nil {
			fmt.Printf("❌ error finding matching templates: %v\n", err)
			return
		}
		for _, name := range matchingNames {
			fmt.Println(name)
		}
	}
}

// Find Matching Template
func findMatchingTemplates(dirPath string, keyword string) ([]string, error) {
	var matchingNames []string

	// Read all files in the directory.
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %w", err)
	}

	// Iterate over the files.
	for _, file := range files {
		// Construct the full file path.
		filePath := filepath.Join(dirPath, file.Name())

		// Check if it's a file and has a .json extension.
		if file.Mode().IsRegular() && strings.ToLower(filepath.Ext(file.Name())) == ".json" {
			// Read the file content.
			fileContent, err := ioutil.ReadFile(filePath)
			if err != nil {
				fmt.Printf("error reading file %s: %v\n", filePath, err)
				// IMPORTANT:  Handle the error.  Here, we print an error and *continue*
				// to the next file.  You might want to return an error, depending
				// on your needs.  Returning here would stop processing *all* files.
				continue // Skip to the next file.
			}

			// Unmarshal the JSON data into the TemplateData struct.
			var data TemplateData
			err = json.Unmarshal(fileContent, &data)
			if err != nil {
				fmt.Printf("error unmarshalling JSON from %s: %v\n", filePath, err)
				continue // Skip to the next file.  Handle the error as appropriate.
			}

			// Check if the keyword is found in either the "name" or "description" (case-insensitive).
			keywordLower := strings.ToLower(keyword)
			nameLower := strings.ToLower(data.Name)
			descriptionLower := strings.ToLower(data.Description)

			if strings.Contains(nameLower, keywordLower) || strings.Contains(descriptionLower, keywordLower) {
				matchingNames = append(matchingNames, data.Name)
			}
		}
	}
	return matchingNames, nil
}