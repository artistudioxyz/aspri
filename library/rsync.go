package library

import (
	"encoding/json"
	"fmt"
	"os"
)

type RsyncConfig struct {
	Flags       string         `json:"flags"`
	Source      RsyncDirectory `json:"source"`
	Destination RsyncDirectory `json:"destination"`
	Excludes    []string       `json:"excludes"`
}

type RsyncDirectory struct {
	Remote string `json:"remote"`
	Path   string `json:"path"`
}

// Initiate File Function
func InitiateRsyncFunction(flags Flag) {
	/** Minify Files in Path .js and .css */
	if *flags.Rsync {
		Rsync()
	}
}

// Rsync command
func Rsync() {
	// Read the JSON file
	file, err := os.Open("rsync.json")
	if err != nil {
		fmt.Println("❌ Error opening JSON file:", err)
		return
	}
	defer file.Close()

	// Parse the JSON data into a RsyncConfig struct
	var config RsyncConfig
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("❌ Error decoding JSON:", err)
		return
	}

	// Check Rsync Remote Configuration.
	checkRemoteConfig := func(directory RsyncDirectory) string {
		if directory.Remote == "" {
			return fmt.Sprintf(`"%s"`, directory.Path)
		} else {
			return fmt.Sprintf(`%s:"%s"`, directory.Remote, directory.Path)
		}
	}

	// Build the rsync command as a string
	rsyncCommand := "rsync " + config.Flags
	for _, exclude := range config.Excludes {
		rsyncCommand += " --exclude=" + exclude
	}
	rsyncCommand += " " + checkRemoteConfig(config.Source)
	rsyncCommand += " " + checkRemoteConfig(config.Destination)

	// Define the shell script content
	scriptContent := "#!/bin/bash\n" + rsyncCommand

	// Create and write the script file
	scriptFileName := "rsync.sh"
	scriptFile, err := os.Create(scriptFileName)
	if err != nil {
		fmt.Println("❌ Error creating script file:", err)
		return
	}
	defer scriptFile.Close()
	_, err = scriptFile.WriteString(scriptContent)
	if err != nil {
		fmt.Println("❌ Error writing to script file:", err)
		return
	}

	// Make the script file executable
	os.Chmod(scriptFileName, 0755)

	// Print a message with the script file location
	fmt.Println("✅ Rsync command has been generated:", rsyncCommand)
	fmt.Println("✅ Shell script has been generated:", scriptFileName)
}
