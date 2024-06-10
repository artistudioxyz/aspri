package library

import (
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
