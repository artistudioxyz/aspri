package library

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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
		RemoveDirectoriesOlderThan(*flags.Path, *flags.Days, *flags.Level, *flags.DryRun)
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

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("❌ Error:", err)
			return err
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
func RemoveDirectoriesOlderThan(path string, retentionDays int, level int, dryrun bool) error {
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
