package library

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// Initiate Directory Function
func InitiateDirectoryFunction(flags Flag) {
	/** Directory Stats */
	if *flags.Dir && *flags.Stats {
		DirectoryStats(*flags.Path, true)
	}
	/** remove Directories Older than days matching regex */
	if *flags.Dir && *flags.Remove && *flags.OlderThan && *flags.Days > 0 {
		RemoveDirectoriesOlderThan(*flags.Path, *flags.Days, *flags.Level, *flags.Exclude, *flags.DryRun)
	}
	/** Normalize Directories Name */
	if *flags.Dir && *flags.Standardize {
		fmt.Println("📁 Standardize Directory Name:", *flags.Path)
		StandardizeDirectoryNameLoop(*flags.Path)
	}
}

// Create a helper function to determine the depth of a directory.
func GetDepth(path string, dirPath string) int {
	rel, err := filepath.Rel(path, dirPath)
	if err != nil {
		return -1 // Error in determining depth.
	}
	return len(strings.Split(rel, "/")) - 1 // Return the depth.
}

/** Directory Stats */
func DirectoryStats(path string, print bool) (int, int64, int64, map[string]int, int, int) {
	if path == "" {
		CurrentDirectory, _ := os.Getwd()
		path = CurrentDirectory
	}

	var count int
	var totalSize int64
	extCount := make(map[string]int)
	var lineCount int
	var wordCount int

	excludedDirs := []string{".stversions", ".obsidian"}

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("❌ Error:", err)
			return err
		}

		// Skip the excluded directories
		for _, excludedDir := range excludedDirs {
			if info.IsDir() && strings.Contains(path, excludedDir) {
				return filepath.SkipDir
			}
		}

		if !info.IsDir() {
			count++
			totalSize += info.Size()

			ext := filepath.Ext(path)
			ext = strings.ToLower(ext)
			extCount[ext]++

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				lineCount++
				wordCount += len(strings.Fields(scanner.Text()))
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("❌ Error:", err)
		return 0, 0, 0, nil, 0, 0
	}

	var averageSize int64
	if count > 0 {
		averageSize = totalSize / int64(count)
	}

	if print {
		currentTime := time.Now()
		fmt.Println("🗓️ Generated at : ", currentTime.String())
		fmt.Println("📈 Total Files:", count)
		fmt.Println("📊 Total Size:", totalSize)
		fmt.Println("💽 Average Size:", averageSize)
		fmt.Println("📝 Total Lines:", lineCount)
		fmt.Println("💬 Total Words:", wordCount)
		fmt.Println("🏺 No of files by extensions :")
		for ext, count := range extCount {
			fmt.Println(" 📟", ext, ":", count)
		}
	}

	return count, totalSize, averageSize, extCount, lineCount, wordCount
}

// Remove directory older than.
func RemoveDirectoriesOlderThan(path string, retentionDays int, level int, exclude []string, dryrun bool) error {
	if path == "" {
		// If path is empty, use the current working directory.
		currentDir, err := os.Getwd()
		if err != nil {
			return err
		}
		path = currentDir
	}

	// Get the current time.
	currentTime := time.Now()

	// Calculate the cutoff date for retention.
	cutoffDate := currentTime.AddDate(0, 0, -retentionDays)

	// Walk through the directories in the provided path.
	err := filepath.Walk(path, func(dirPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the directory should be ignored
		for _, excludeDir := range exclude {
			if strings.Contains(dirPath, excludeDir) {
				return nil
			}
		}

		// Check if the directory is within the specified depth.
		if GetDepth(path, dirPath) <= level {
			// Check if the directory is older than the retention cutoff date.
			if info.IsDir() && info.ModTime().Before(cutoffDate) {
				if !dryrun {
					fmt.Println("✅ Successfully remove directories older than", retentionDays, "days in", dirPath)
					err := os.RemoveAll(dirPath)
					if err != nil {
						return err
					}
				} else {
					fmt.Println("✅ Dry run, will remove", dirPath)
				}
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// Standardize Directory Name
func StandardizeDirectoryName(rootPath string) error {
	// Define a regular expression to match emojis
	emojiRegex := regexp.MustCompile(`[\p{So}\p{Sk}]`)

	// Walk through the directory tree
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the path is a directory
		if info.IsDir() {
			// Get the directory name
			dirName := filepath.Base(path)

			// Check if the directory name contains an emoji
			if emojiRegex.MatchString(dirName) {
				// Remove emojis from the directory name
				dirName = emojiRegex.ReplaceAllString(dirName, "")
				fmt.Printf("Renaming directory: %s -> %s\n", filepath.Join(filepath.Dir(path), filepath.FromSlash(dirName)), filepath.Join(filepath.Dir(path), dirName))
				// Rename the directory
				if err := os.Rename(path, filepath.Join(filepath.Dir(path), dirName)); err != nil {
					return err
				}
				return &CustomError{message: "Directory name contains an space"}
			}

			// Replace spaces with underscores in the directory name
			dirName = strings.ReplaceAll(dirName, " ", "_")
			if dirName != filepath.Base(path) {
				fmt.Printf("Renaming directory: %s -> %s\n", filepath.Join(filepath.Dir(path), filepath.FromSlash(dirName)), filepath.Join(filepath.Dir(path), dirName))
				// Rename the directory
				if err := os.Rename(path, filepath.Join(filepath.Dir(path), dirName)); err != nil {
					return err
				}
				return &CustomError{message: "Directory name contains an space"}
			}
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Loop through the directory tree and standardize directory names
func StandardizeDirectoryNameLoop(rootPath string) {
	err := StandardizeDirectoryName(rootPath)

	// Check if the error is an instance of CustomError
	if _, ok := err.(*CustomError); ok {
		StandardizeDirectoryNameLoop(rootPath)
	} else {
		fmt.Println("Regular Error:", err)
	}
}
