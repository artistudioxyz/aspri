package wordpress

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

/** Plugin Type */
type Plugin struct {
	Name     string
	FilePath string
	Version  string
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

	/** Read File */
	content, err := ioutil.ReadFile(plugin.FilePath)
	if err != nil {
		log.Fatal(err)
	}

	/** Read Comment Block */
	musComp := regexp.MustCompile("(?s)//.*?\n|/\\*.*?\\*/")
	comments := strings.Split(musComp.FindString(string(content)), "\n")
	for _, s := range comments {
		if strings.Contains(s, "Version:") {
			plugin.Version = strings.Replace(s, "Version:", "", -1)
			plugin.Version = strings.Replace(plugin.Version, "*", "", -1)
			plugin.Version = strings.Replace(plugin.Version, " ", "", -1)
		}
	}

	fmt.Println(plugin)
}
