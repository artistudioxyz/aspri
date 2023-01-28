package library

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

/** Initiate Syncthing Function */
func InitiateSyncthingFunction(flags Flag) {
	/** remove Sync Conflict Files older Than x days */
	if *flags.Syncthing && *flags.RemoveConflicts && *flags.Days > 0 {
		removeSyncConflictFiles(*flags.Path, *flags.Days, *flags.DryRun)
	}
}

/** remove Sync Conflict Files older Than x days */
func removeSyncConflictFiles(path string, retentionDays int, dryRun bool) error {
	if path == "" {
		CurrentDirectory, _ := os.Getwd()
		path = CurrentDirectory
	}

	currentTime := time.Now()
	dateFormat := "20060102"
	dateRegex := regexp.MustCompile(`sync-conflict-(\d{8})-`)
	return filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		fileName := strings.TrimSuffix(info.Name(), filepath.Ext(filePath))
		if match := dateRegex.FindStringSubmatch(fileName); len(match) > 1 {
			fileDate, _ := time.Parse(dateFormat, match[1])
			if !fileDate.Add(time.Duration(retentionDays) * 24 * time.Hour).After(currentTime) {
				if dryRun {
					fmt.Println("✅ Dry run, will remove", filePath)
				} else {
					os.Remove(filePath)
					fmt.Println("✅ Successfully remove conflict ", filePath)
				}
			}
		}
		return nil
	})
}
