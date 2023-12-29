package library

import (
	"fmt"
	"regexp"
	"strings"
)

/** Initiate Miscellaneous Function */
func InitiateMiscellaneousFunction(flags Flag) {
	/** Self Update */
	if *flags.SelfUpdate {
		fmt.Println("✅ Doing self update")
		cmd := [...]string{"bash", "-c", "go get github.com/artistudioxyz/aspri"}
		ExecCommand(cmd[:]...)
	}
}

/** Slugify function */
func Slugify(s string) string {
	// Convert the string to lowercase
	s = strings.ToLower(s)

	// Replace non-alphanumeric characters with spaces, except for forward slashes and periods
	re := regexp.MustCompile(`[^a-z0-9/.]+`)
	s = re.ReplaceAllString(s, " ")

	// Replace all " / " with "/"
	s = strings.ReplaceAll(s, " /", "/")

	// Trim leading and trailing spaces
	s = strings.TrimSpace(s)

	// Replace spaces with hyphens
	s = strings.ReplaceAll(s, " ", "-")

	// Remove leading and trailing hyphens
	s = strings.Trim(s, "-")

	// Replace consecutive forward slashes with single forward slashes
	re = regexp.MustCompile(`/+`)
	s = re.ReplaceAllString(s, "/")

	return s
}
