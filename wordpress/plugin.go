package wordpress

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/** GetPluginInformation */
func GetPluginInformation(path string) WPProject {
	plugin := WPProject{}
	plugin.Path.Directory = path

	PathArray := strings.Split(path, string(filepath.Separator))
	plugin.Name = PathArray[len(PathArray)-1]
	FileName := fmt.Sprintf("%s.php", PathArray[len(PathArray)-1])
	plugin.Path.File = strings.Join(append(PathArray, FileName), string(filepath.Separator))

	// Check if file exists
	if _, err := os.Stat(plugin.Path.File); err != nil {
		if os.IsNotExist(err) {
			PathArray := strings.Split(path, string(filepath.Separator))
			plugin.Name = PathArray[len(PathArray)-2]
			FileName := fmt.Sprintf("%s.php", PathArray[len(PathArray)-2])
			plugin.Path.File = strings.Join(append(PathArray, FileName), string(filepath.Separator))
		}
	}

	return ReadCommentBlock(plugin)
}
