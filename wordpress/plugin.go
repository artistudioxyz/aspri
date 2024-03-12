package wordpress

import (
	"fmt"
	"strings"
)

/** GetPluginInformation */
func GetPluginInformation(path string) WPProject {
	plugin := WPProject{}
	plugin.Path.Directory = path

	PathArray := strings.Split(path, "/")
	plugin.Name = PathArray[len(PathArray)-1]
	FileName := fmt.Sprintf("%s.php", PathArray[len(PathArray)-1])
	PathArray = append(PathArray, FileName)
	plugin.Path.File = strings.Join(PathArray, "/")

	return ReadCommentBlock(plugin)
}
