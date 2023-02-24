package library

import (
	"bufio"
	"fmt"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/js"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

/** Initiate Miscellaneous Function */
func InitiateMiscellaneousFunction(flags Flag) {
	/** Minify Files in Path .js and .css */
	if *flags.Minify {
		minifyFiles(*flags.Path)
	}
	/** Count Files Containing Text */
	if *flags.File && *flags.Count && *flags.Text != "" {
		count := CountFilesContainingText(*flags.Path, *flags.Text)
		fmt.Println("🐙 There are", count, "files containing", *flags.Text)
	}
	/** Directory Stats */
	if *flags.Dir && *flags.Stats {
		DirectoryStats(*flags.Path, true)
	}
	/** removeFilesExceptExtensions */
	if *flags.File && *flags.Remove && len(*flags.Ext) > 0 {
		RemoveFilesExceptExtensions(*flags.Path, *flags.Ext, *flags.Except)
	}
	/** remove Files Older than days matching regex */
	if *flags.File && *flags.Remove && *flags.OlderThan && *flags.Days > 0 && *flags.Regex != "" {
		RemoveFilesOlderThan(*flags.Path, *flags.Regex, *flags.Days, *flags.DryRun)
	}
	/** Delete Directory or Files in Path Matching Filename */
	if *flags.Dir && *flags.Remove && len(*flags.Dirname) > 0 {
		DeleteDirectoriesorFilesinPath(*flags.Path, *flags.Dirname, *flags.Filename)
	}
	if *flags.File && *flags.Remove && len(*flags.Filename) > 0 {
		DeleteDirectoriesorFilesinPath(*flags.Path, *flags.Dirname, *flags.Filename)
	}
	/** Search and Replace */
	if *flags.SearchandReplace && *flags.From != "" && *flags.To != "" {
		SearchandReplace(*flags.Path, *flags.From, *flags.To)
	}
	/** Self Update */
	if *flags.SelfUpdate {
		fmt.Println("✅ Doing self update")
		cmd := [...]string{"bash", "-c", "go get github.com/artistudioxyz/aspri"}
		ExecCommand(cmd[:]...)
	}
}

/** Minify Files in Path .js and .css */
func minifyFiles(path string) {
	if path == "" {
		CurrentDirectory, _ := os.Getwd()
		path = CurrentDirectory
	}

	m := minify.New()
	m.AddFunc("text/javascript", js.Minify)
	m.AddFunc("text/css", css.Minify)
	filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("❌", err)
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(filePath) == ".js" || filepath.Ext(filePath) == ".css" {
			// Open the file
			file, err := os.Open(filePath)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			// Minify the file
			var contentType string
			if filepath.Ext(filePath) == ".js" {
				contentType = "text/javascript"
			} else {
				contentType = "text/css"
			}

			// read the file
			bs, err := ioutil.ReadAll(file)
			if err != nil {
				panic(err)
			}

			// minify the content
			minifiedContent, err := m.String(contentType, string(bs))
			if err != nil {
				panic(err)
			}

			// write the minified content to the file
			err = ioutil.WriteFile(filePath, []byte(minifiedContent), 0644)
			if err != nil {
				panic(err)
			}
		}
		return nil
	})

	fmt.Println("✅ Successfully minify files in", path)
}

/** Count Files Containing Text */
func CountFilesContainingText(path string, text string) int {
	if path == "" {
		CurrentDirectory, _ := os.Getwd()
		path = CurrentDirectory
	}

	var count int

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("❌ Error:", err)
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				if strings.Contains(scanner.Text(), text) {
					count++
					break
				}
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("❌ Error:", err)
		return 0
	}

	return count
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

/** Remove Files Except Specified Extensions */
func RemoveFilesExceptExtensions(root string, allowedExtensions []string, exception []string) error {
	if root == "" {
		CurrentDirectory, _ := os.Getwd()
		root = CurrentDirectory
	}

	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("❌ ", err)
		}
		if !info.IsDir() {
			ext := filepath.Ext(info.Name())
			if !SliceContainsString(allowedExtensions, ext) && !SliceContainsString(exception, info.Name()) {
				err := os.Remove(path)
				if err != nil {
					return err
				}
				fmt.Println("✅ Successfully remove files except extensions", allowedExtensions, "in", info.Name())
			}
		}
		return nil
	})
}

/** Remove Files Older Than */
func RemoveFilesOlderThan(path string, pattern string, retentionDays int, dryrun bool) error {
	if path == "" {
		CurrentDirectory, _ := os.Getwd()
		path = CurrentDirectory
	}

	currentTime := time.Now()
	return filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("❌ ", err)
		}
		if info.IsDir() {
			return nil
		}
		if !info.ModTime().Add(time.Duration(retentionDays) * 24 * time.Hour).After(currentTime) {
			if matched, _ := filepath.Match(pattern, info.Name()); matched {
				if dryrun {
					fmt.Println("✅ Dry run, will remove", filePath)
				} else {
					os.Remove(filePath)
					fmt.Println("✅ Successfully remove files older than", retentionDays, "days in", filePath)
				}
			}
		}
		return nil
	})
}

/** Delete Directory or Files in Path Matching Filename */
func DeleteDirectoriesorFilesinPath(root string, dirnames []string, filenames []string) error {
	if root == "" {
		CurrentDirectory, _ := os.Getwd()
		root = CurrentDirectory
	}

	// Walk through the directory tree
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("❌ ", err)
			return nil
		}

		// If the path is a directory and it has the correct name, delete it
		if SliceContainsString(dirnames, info.Name()) || SliceContainsString(filenames, info.Name()) {
			err = os.RemoveAll(path)
			if err != nil {
				fmt.Println("❌ ", err)
				return nil
			}
			if info.IsDir() {
				fmt.Println("✅ Successfully remove directories nested by name", info.Name(), "in", root)
			} else {
				fmt.Println("✅ Successfully remove files nested by filename", info.Name(), "in", root)
			}
		} else if info.IsDir() {
			// Check if the directory is empty
			f, err := os.Open(path)
			if err != nil {
				fmt.Println("❌ ", err)
				return nil
			}
			defer f.Close()
			_, err = f.Readdirnames(1)
			if err == io.EOF {
				// Directory is empty, so delete it
				os.Remove(path)
			}
		}

		return nil
	})
}

/** Search and Replace in Directory or File */
func SearchandReplace(path string, from string, to string) {
	if path == "" {
		CurrentDirectory, _ := os.Getwd()
		path = CurrentDirectory
	}

	/** Search and Replace */
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			newData := strings.Replace(string(data), from, to, -1)
			err = ioutil.WriteFile(path, []byte(newData), 0644)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("❌ ", err)
	} else {
		fmt.Println("✅ Success Search and Replace", from, "to", to, "in", path)
	}
}
