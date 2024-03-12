package wordpress

import (
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

	return ReadCommentBlock(theme)
}
