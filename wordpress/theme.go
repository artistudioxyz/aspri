package wordpress

import (
	"path/filepath"
	"strings"
)

/** GetThemeInformation */
func GetThemeInformation(path string) WPProject {
	theme := WPProject{}
	theme.Path.Directory = path

	PathArray := strings.Split(path, string(filepath.Separator))
	theme.Name = PathArray[len(PathArray)-1]
	PathArray = append(PathArray, "style.css")
	theme.Path.File = strings.Join(PathArray, string(filepath.Separator))

	return ReadCommentBlock(theme)
}
