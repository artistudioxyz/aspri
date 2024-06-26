package library

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Initiate Markdown Function
func InitiateMarkdownFunction(flags Flag) {
	/** Remove Link */
	if *flags.Markdown && *flags.RemoveLink {
		if *flags.Path == "" {
			CurrentDirectory, _ := os.Getwd()
			*flags.Path = CurrentDirectory
		}
		content := ReadFile(*flags.Path)
		WriteFile(*flags.Path, MarkdownRemoveLink(string(content)))
		fmt.Println("✅ Link removed")
	}
	/** Remove Link */
	if *flags.Markdown && *flags.Tree {
		fileTree := MarkdownGenerateFileTree(*flags.Path, *flags.Filename)
		fmt.Println(fileTree)
	}
	// Extract Heading Number
	if *flags.Markdown && *flags.Heading != "" && strings.HasPrefix(*flags.Heading, "#") {
		headings, _ := ExtractHeadings(*flags.Path, *flags.Heading)
		fmt.Println(headings)
	}
	// Extract Markdown content
	if *flags.Markdown && *flags.Heading != "" {
		content, _ := ExtractContentByHeading(*flags.Path, *flags.Heading)
		fmt.Println(content)
	}
}

// Remove Link from Markdown File
func MarkdownRemoveLink(markdown string) string {
	// Use the regular expression to search for all links
	pattern := regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)
	links := pattern.FindAllStringSubmatch(markdown, -1)

	// Replace each link with its label (the text between the square brackets)
	for _, link := range links {
		markdown = strings.Replace(markdown, link[0], link[1], -1)
	}

	return markdown
}

// Generate File Tree from Directory
func MarkdownGenerateFileTree(path string, ignore []string) string {
	if path == "" {
		CurrentDirectory, _ := os.Getwd()
		path = CurrentDirectory
	}

	// Ignore the files and directories
	ignore = append(ignore, ".git", ".github", ".vscode", ".idea", ".obsidian", ".gitignore", ".gitkeep", ".DS_Store")

	// Read the directory
	files, err := os.ReadDir(path)
	if err != nil {
		return ""
	}

	tree := ""
	indent := strings.Repeat("    ", strings.Count(path, string(os.PathSeparator)))

	for _, file := range files {
		if SliceContainsString(ignore, file.Name()) {
			continue
		}

		if file.IsDir() {
			subtree := MarkdownGenerateFileTree(filepath.Join(path, file.Name()), ignore)
			tree += fmt.Sprintf("%s- %s\n", indent, file.Name())
			tree += subtree
		} else {
			link := filepath.Join(strings.Split(path, string(os.PathSeparator))...)
			slugifyPath := Slugify(link + string(os.PathSeparator) + file.Name())
			tree += fmt.Sprintf("%s- [%s](%s)\n", indent, file.Name(), slugifyPath)
		}
	}

	return tree
}

// Extract heading
func ExtractHeadings(filePath, heading string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var headings []string
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(fmt.Sprintf(`^%s\s+(.*)$`, regexp.QuoteMeta(heading)))

	for scanner.Scan() {
		line := scanner.Text()
		match := re.FindStringSubmatch(line)
		if len(match) > 1 {
			headings = append(headings, match[1])
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return headings, nil
}

// Extract markdown content by heading
func ExtractContentByHeading(filePath, heading string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "", err
	}
	defer file.Close()

	var content strings.Builder
	scanner := bufio.NewScanner(file)
	found := false

	for scanner.Scan() {
		line := scanner.Text()

		// Check for headings
		if strings.HasPrefix(line, "#") {
			// Trim leading spaces and hash symbols
			currentHeading := strings.TrimLeft(line, "# ")
			if currentHeading == heading {
				found = true
				continue // Skip the heading line itself
			}
		}

		if found {
			// Append content until next heading or end of file
			if strings.HasPrefix(line, "#") {
				break // Stop if a new heading is found
			}
			fmt.Fprintln(&content, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return content.String(), nil
}
