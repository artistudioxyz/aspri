package wordpress

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

/** Plugin Type */
type Plugin struct {
	Name     string
	FilePath string
	Content  string
}

/* WP Plugin Check */
func WPPluginBuildCheck(path string, production bool) {
	fmt.Println("Check Plugin")
	fmt.Println(path)

	/** Get Plugin Information */
	PathArray := strings.Split(path, "/")
	plugin := Plugin{}
	plugin.Name = PathArray[len(PathArray)-1]
	PathArray = append(PathArray, fmt.Sprintf("%s.php", PathArray[len(PathArray)-1]))
	plugin.FilePath = strings.Join(PathArray, "/")

	fmt.Println(plugin)

	/** Read File */
	content, err := ioutil.ReadFile(plugin.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}
