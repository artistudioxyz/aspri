package wordpress

import (
	"fmt"
	"os"
	"strings"
)

/** GetThemeInformation */
func GetThemeInformation(path string) WPProject {
	theme := WPProject{}
	theme.Path.Directory = path

	PathArray := strings.Split(path, "/")
	theme.Name = PathArray[len(PathArray)-1]
	PathArray = append(PathArray, "style.css")
	theme.Path.File = strings.Join(PathArray, "/")

	return theme
}

/* WP Theme Check */
func WPThemeBuildCheck(path string) {
	fmt.Println("🔍 Check Theme")
	if path == "" {
		CurrentDirectory, _ := os.Getwd()
		path = CurrentDirectory
	}
	theme := GetThemeInformation(path)
	CheckProjectVersion(theme)
}
