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

// PHPCS Config
type PHPCSConfig struct {
	Phpcs string `json:"phpcs"`
}

/** Initiate PHPCS Function */
func InitiatePHPCSFunction(flags Flag) {
	/** PHPCS Install Ruleset */
	if *flags.PHPCS && *flags.Install {
		phpCSInstallRuleset()
	}
}

// Get PHPCSConfig
func phpCSGetConfig() (PHPCSConfig, error) {
	// Get current directory
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("❌", err)
		os.Exit(1)
	}

	// Define the path to the config file.
	configPath := dir + "/config.json"

	// Read the contents of the config file.
	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println("❌", err)
		return PHPCSConfig{}, err
	}

	// Parse the JSON data into a Config struct.
	var config PHPCSConfig
	err = json.Unmarshal(configData, &config)
	if err != nil {
		fmt.Println("❌", err)
		return PHPCSConfig{}, err
	}

	return config, nil
}

// Detect Standard
func phpCSDetectStandard() ([]string, error) {
	var standards []string
	var standardsDirectory []string

	// Get current directory
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return standards, err
	}

	// Set standards.json file path
	standardsJSON := dir + "/standards.json"

	// Read standards.json file
	bytes, err := ioutil.ReadFile(standardsJSON)
	if err != nil {
		return standards, err
	}

	// Unmarshal standards.json data into slice of strings
	err = json.Unmarshal(bytes, &standardsDirectory)
	if err != nil {
		return standards, err
	}

	// Add standards directory to standards slice
	standardsDirectory = append(standardsDirectory, dir+"/standards")

	for _, standardDirectory := range standardsDirectory {
		// List all subdirectories in standards directory
		subdirectories, err := ioutil.ReadDir(standardDirectory)
		if err != nil {
			return standards, err
		}

		// Add subdirectories to standards slice
		for _, subdirectory := range subdirectories {
			// Check ruleset.xml file exists in subdirectory
			FileExist, _ := FileExistsInPath("ruleset.xml", standardDirectory+"/"+subdirectory.Name())
			if subdirectory.IsDir() && FileExist {
				standards = append(standards, standardDirectory+"/"+subdirectory.Name())
			}
		}
	}

	return standards, nil
}

/** PHPCS Install Ruleset */
func phpCSInstallRuleset() {
	// Set PHPCS path
	phpcsConfig, err := phpCSGetConfig()
	phpcsPath := "phpcs"
	if err == nil {
		phpcsPath = phpcsConfig.Phpcs
	}

	fmt.Println("🔍 PHPCS path:", phpcsPath)

	// Detect standards
	standards, err := phpCSDetectStandard()
	if err != nil {
		fmt.Println("❌", err)
		os.Exit(1)
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
	cmd := exec.Command(phpcsPath, "--config-set", "installed_paths", `"`+installedPaths+`"`)
	fmt.Println("📟 Execute :", cmd)
	err = cmd.Run()
	if err != nil {
		fmt.Println("❌", err)
		os.Exit(1)
	}
}
