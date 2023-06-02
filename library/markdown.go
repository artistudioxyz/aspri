package library

import (
	"fmt"
	"os"
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
