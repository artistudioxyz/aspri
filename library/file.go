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
	"regexp"
	"strings"
	"time"
)

// Initiate File Function
func InitiateFileFunction(flags Flag) {
	/** Minify Files in Path .js and .css */
	if *flags.Minify {
		minifyFiles(*flags.Path)
	}
	/** Count Files Containing Text */
	if *flags.File && *flags.Count && *flags.Text != "" {
		count := CountFilesContainingText(*flags.Path, *flags.Text, *flags.IgnoreDirs)
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
	if *flags.File && *flags.Remove && *flags.OlderThan && *flags.Days > 0 {
		RemoveFilesOlderThan(*flags.Path, *flags.Regex, *flags.Days, *flags.DryRun)
	}
	/** remove Directories Older than days matching regex */
	if *flags.Dir && *flags.Remove && *flags.OlderThan && *flags.Days > 0 {
		RemoveDirectoriesOlderThan(*flags.Path, *flags.Days, *flags.Level, *flags.DryRun)
	}
	/** Delete Directory or Files in Path Matching Filename */
	if *flags.Dir && *flags.Remove && len(*flags.Dirname) > 0 {
		DeleteDirectoriesorFilesinPath(*flags.Path, *flags.Dirname, *flags.Filename)
	}
	if *flags.File && *flags.Remove && len(*flags.Filename) > 0 {
		DeleteDirectoriesorFilesinPath(*flags.Path, *flags.Dirname, *flags.Filename)
	}
	/** Exctract Links from Directory Path */
	if *flags.ExtractUrl {
		urls, err := ExtractURLsFromDirectoryPath(*flags.Path, *flags.Url)
		if err != nil {
			fmt.Println("❌ Error extracting links:", err)
		}
		for _, url := range urls {
			fmt.Println(url)
		}
	}
	/** Search and Replace */
	if *flags.SearchandReplace && *flags.From != "" && *flags.To != "" {
		SearchandReplace(*flags.Path, *flags.From, *flags.To)
	}
}

/** Extract URLs from File */
func extractURLsFromFile(filePath string, baseURL string) ([]string, error) {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Define a regular expression to match URLs
	urlPattern := `https?://\S+`

	var urls []string

	// Find URLs in the file content
	urlRegex := regexp.MustCompile(urlPattern)
	urlMatches := urlRegex.FindAllString(string(fileContent), -1)

	// Filter URLs based on the baseURL if provided
	if baseURL != "" {
		for _, url := range urlMatches {
			if strings.Contains(url, baseURL) {
				urls = append(urls, url)
			}
		}
	} else {
		urls = append(urls, urlMatches...)
	}

	return urls, nil
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
func CountFilesContainingText(path string, text string, ignoreDirs []string) int {
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

		// Check if the directory should be ignored
		for _, ignoreDir := range ignoreDirs {
			if info.IsDir() && info.Name() == ignoreDir {
				return filepath.SkipDir
			}
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
			// Match file name if pattern is not empty
			matched := true
			if pattern != "" {
				matched, _ = filepath.Match(pattern, info.Name())
			}

			// Remove file if matched
			if matched {
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

	// Create a helper function to determine the depth of a directory.
	getDepth := func(dirPath string) int {
		rel, err := filepath.Rel(path, dirPath)
		if err != nil {
			return -1 // Error in determining depth.
		}
		return len(filepath.SplitList(rel))
	}

	// Walk through the directories in the provided path.
	err := filepath.Walk(path, func(dirPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the directory is within the specified depth.
		if getDepth(dirPath) <= level {
			// Check if the directory is older than the retention cutoff date.
			if info.IsDir() && info.ModTime().Before(cutoffDate) {
				fmt.Printf("Removing directory: %s\n", dirPath)
				if !dryrun {
					err := os.RemoveAll(dirPath)
					if err != nil {
						return err
					}
				} else {
					fmt.Println("(Dry Run: No removal performed)")
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

/** Exctract URLs from Directory Path */
func ExtractURLsFromDirectoryPath(path string, baseURL string) ([]string, error) {
	if path == "" {
		// Use the current directory path if path is not provided
		dir, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		path = dir
	}

	uniqueURLs := make(map[string]struct{}) // Map to store unique URLs

	// Check if the path is a directory
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if !fileInfo.IsDir() {
		return nil, fmt.Errorf("Path is not a directory: %s", path)
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			// If it's a subdirectory, recursively extract URLs
			subpath := filepath.Join(path, file.Name())
			subURLs, err := ExtractURLsFromDirectoryPath(subpath, baseURL)
			if err != nil {
				return nil, err
			}
			for _, url := range subURLs {
				uniqueURLs[url] = struct{}{}
			}
		} else {
			// If it's a file, extract URLs based on the file content
			filePath := filepath.Join(path, file.Name())
			fileURLs, err := extractURLsFromFile(filePath, baseURL)
			if err != nil {
				return nil, err
			}
			for _, url := range fileURLs {
				uniqueURLs[url] = struct{}{}
			}
		}
	}

	// Convert unique URLs from the map to a slice
	var urls []string
	for url := range uniqueURLs {
		// Cleaned unwanted symbols from the URL
		symbolPattern := `[^\w://.]`
		regex := regexp.MustCompile(symbolPattern)
		cleaned := regex.ReplaceAllString(url, "")

		// Add the URL to the slice
		urls = append(urls, cleaned)
	}

	return urls, nil
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
