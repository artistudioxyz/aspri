package library

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

/** Initiate PHPCS Function */
func InitiatePHPCSFunction(flags Flag) {
	/** PHPCS Install Ruleset */
	if *flags.PHPCS && *flags.Install {
		phpCSInstallRuleset()
	}
}

/** PHPCS Install Ruleset */
func phpCSInstallRuleset() {
	// Get current directory
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("❌", err)
		os.Exit(1)
	}

	// Set standards.json file path
	standardsJSON := dir + "/standards.json"

	// Set standards directory path
	standardsDirectory := dir + "/standards"

	// Read standards.json file
	bytes, err := ioutil.ReadFile(standardsJSON)
	if err != nil {
		fmt.Println("❌", err)
		os.Exit(1)
	}

	// Unmarshal standards.json data into slice of strings
	var standards []string
	err = json.Unmarshal(bytes, &standards)
	if err != nil {
		fmt.Println("❌", err)
		os.Exit(1)
	}

	// List all subdirectories in standards directory
	subdirectories, err := ioutil.ReadDir(standardsDirectory)
	if err != nil {
		fmt.Println("❌", err)
		os.Exit(1)
	}

	// Add subdirectories to standards slice
	for _, subdirectory := range subdirectories {
		if subdirectory.IsDir() {
			standards = append(standards, standardsDirectory+"/"+subdirectory.Name())
		}
	}

	// Detected standards
	fmt.Println("🔍 Detected standards:", strings.Join(standards, ","))

	// Join standards slice into a single string
	installedPaths := ""
	for i, path := range standards {
		installedPaths += path
		if i < len(standards)-1 {
			installedPaths += ","
		}
	}

	// Execute PHPCS command to set installed paths
	cmd := exec.Command("phpcs", "--config-set", "installed_paths", installedPaths)
	err = cmd.Run()
	if err != nil {
		fmt.Println("❌", err)
		os.Exit(1)
	}
}
