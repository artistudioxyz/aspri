package wordpress

import (
	"fmt"
	"os"
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

	return plugin
}

/* WP Plugin Check */
func WPPluginBuildCheck(path string) {
	fmt.Println("🔍 Check Theme")
	if path == "" {
		CurrentDirectory, _ := os.Getwd()
		path = CurrentDirectory
	}
	plugin := GetPluginInformation(path)
	CheckProjectVersion(plugin)
}
