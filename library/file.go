package library

import (
	"bufio"
	"fmt"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/js"
	"io"
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
		count := CountFilesContainingText(*flags.Path, *flags.Text, *flags.Exclude)
		fmt.Println("🐙 There are", count, "files containing", *flags.Text)
	}
	// Find files younger than days matching regex
	if *flags.File && *flags.Find && *flags.YoungerThan && *flags.Days > 0 {
		files, err := FindFilesByAge(*flags.Path, *flags.Regex, *flags.Days, *flags.Exclude, false)
		if err != nil {
			fmt.Println("❌ ", err)
		} else {
			for _, file := range files {
				fmt.Println(file)
			}
		}
	}
	// Find files older than days matching regex
	if *flags.File && *flags.Find && *flags.OlderThan && *flags.Days > 0 {
		files, err := FindFilesByAge(*flags.Path, *flags.Regex, *flags.Days, *flags.Exclude, true)
		if err != nil {
			fmt.Println("❌ ", err)
		} else {
			for _, file := range files {
				fmt.Println(file)
			}
		}
	}
	// Find files between dates matching regex
	if *flags.File && *flags.Find && *flags.Between && *flags.Start != "" && *flags.End != "" {
		files, err := FindFilesBetweenDates(*flags.Path, *flags.Regex, *flags.Start, *flags.End, *flags.Exclude)
		if err != nil {
			fmt.Println("❌ ", err)
		} else {
			for _, file := range files {
				fmt.Println(file)
			}
		}
	}
	/** removeFilesExceptExtensions */
	if *flags.File && *flags.Remove && len(*flags.Ext) > 0 {
		RemoveFilesExceptExtensions(*flags.Path, *flags.Ext, *flags.Except)
	}
	/** remove Files Older than days matching regex */
	if *flags.File && *flags.Remove && *flags.OlderThan && *flags.Days > 0 {
		files, err := FindFilesByAge(*flags.Path, *flags.Regex, *flags.Days, *flags.Exclude, true)
		if err != nil {
			panic(err)
		}

		err = RemoveFilesOlderThan(files, *flags.DryRun)
		if err != nil {
			panic(err)
		}
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
	/** Search and Replace in File or Directory */
	if *flags.SearchandReplace && *flags.From != "" && len(*flags.Filename) > 0 {
		SearchandReplaceFiles(*flags.Filename, *flags.From, *flags.To)
	} else if *flags.SearchandReplace && *flags.From != "" && len(*flags.Filename) == 0 {
		SearchandReplaceDirectory(*flags.Path, *flags.From, *flags.To, -1)
	}
}

// File Exist in Path.
func FileExistsInPath(filePath, directoryPath string) (bool, error) {
	// Construct the full path to the file.
	fullPath := filepath.Join(directoryPath, filePath)

	// Check if the file exists.
	_, err := os.Stat(fullPath)
	if err == nil {
		// File exists.
		return true, nil
	} else if os.IsNotExist(err) {
		// File does not exist.
		return false, nil
	} else {
		// An error occurred while accessing the file.
		return false, err
	}
}

/** Extract URLs from File */
func extractURLsFromFile(filePath string, baseURL string) ([]string, error) {
	fileContent, err := os.ReadFile(filePath)
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
			bs, err := io.ReadAll(file)
			if err != nil {
				panic(err)
			}

			// minify the content
			minifiedContent, err := m.String(contentType, string(bs))
			if err != nil {
				panic(err)
			}

			// write the minified content to the file
			err = os.WriteFile(filePath, []byte(minifiedContent), 0644)
			if err != nil {
				panic(err)
			}
		}
		return nil
	})

	fmt.Println("✅ Successfully minify files in", path)
}

/** Count Files Containing Text */
func CountFilesContainingText(path string, text string, exclude []string) int {
	var count int

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("❌ Error:", err)
			return err
		}

		// Check if the directory should be ignored
		for _, excludeDir := range exclude {
			if strings.Contains(path, excludeDir) {
				return nil
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

// Find files by age
func FindFilesByAge(path string, pattern string, retentionDays int, exclude []string, older bool) ([]string, error) {
	var files []string

	// Calculate the cutoff time
	cutoff := time.Now().AddDate(0, 0, -retentionDays)

	// Walk through the directory
	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file matches the pattern
		if !info.IsDir() && strings.Contains(info.Name(), pattern) {
			// Check if the file meets the age criteria
			fileMatchesAgeCriteria := (older && info.ModTime().Before(cutoff)) || (!older && info.ModTime().After(cutoff))

			if fileMatchesAgeCriteria {
				// Check if the file is in the exclude list
				excluded := false
				for _, excludeFile := range exclude {
					if strings.Contains(filePath, excludeFile) {
						excluded = true
						break
					}
				}

				// If not excluded, add to the list of files
				if !excluded {
					files = append(files, filePath)
				}
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

// Find files between dates matching regex
func FindFilesBetweenDates(path string, pattern string, start string, end string, exclude []string) ([]string, error) {
	var files []string

	// Parse the start and end dates
	startDate, err := time.Parse("2006-01-02", start)
	if err != nil {
		return nil, fmt.Errorf("invalid start date: %w", err)
	}
	endDate, err := time.Parse("2006-01-02", end)
	if err != nil {
		return nil, fmt.Errorf("invalid end date: %w", err)
	}

	// Walk through the directory
	err = filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file matches the pattern
		if !info.IsDir() && strings.Contains(info.Name(), pattern) {
			// Check if the file was modified between the start and end dates
			fileDate := info.ModTime()
			if fileDate.After(startDate) && fileDate.Before(endDate) {
				// Check if the file is in the exclude list
				excluded := false
				for _, excludeFile := range exclude {
					if strings.Contains(filePath, excludeFile) {
						excluded = true
						break
					}
				}

				// If not excluded, add to the list of files
				if !excluded {
					files = append(files, filePath)
				}
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
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

// Remove Files Older Than
func RemoveFilesOlderThan(files []string, dryrun bool) error {
	for _, file := range files {
		if dryrun {
			fmt.Println("Would remove:", file)
		} else {
			err := os.Remove(file)
			if err != nil {
				return fmt.Errorf("failed to remove file %s: %w", file, err)
			}
			fmt.Println("Removed:", file)
		}
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

	files, err := os.ReadDir(path)
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

/** Search and Replace in File */
func SearchandReplaceFiles(files []string, from string, to string) error {
	for _, filePath := range files {
		// Open the file for reading and writing
		file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
		if err != nil {
			return err
		}
		defer file.Close()

		// Create a scanner to read the file line by line
		scanner := bufio.NewScanner(file)

		// Create a temporary file to store the modified content
		tmpFile, err := os.CreateTemp("", "tmp")
		if err != nil {
			return err
		}
		defer tmpFile.Close()

		// Create a writer to write to the temporary file
		writer := bufio.NewWriter(tmpFile)

		// Iterate over each line in the file
		for scanner.Scan() {
			// Read the line
			line := scanner.Text()

			// Search and replace the string
			modifiedLine := strings.ReplaceAll(line, from, to)

			// Write the modified line to the temporary file
			_, err := writer.WriteString(modifiedLine + "\n")
			if err != nil {
				return err
			}
		}

		// Flush the writer to ensure all buffered data is written to the file
		if err := writer.Flush(); err != nil {
			return err
		}

		// Close the temporary file
		if err := tmpFile.Close(); err != nil {
			return err
		}

		// Remove the original file
		if err := os.Remove(filePath); err != nil {
			return err
		}

		// Rename the temporary file to the original file name
		if err := os.Rename(tmpFile.Name(), filePath); err != nil {
			return err
		}
	}

	return nil
}

/** Search and Replace in Directory */
func SearchandReplaceDirectory(path string, from string, to string, limit int) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			newData := strings.Replace(string(data), from, to, limit)
			err = os.WriteFile(path, []byte(newData), 0644)
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
