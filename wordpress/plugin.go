package wordpress

import (
	"aspri/library"
	"fmt"
	"regexp"
	"strings"
)

/** Path Type */
type WPPluginPath struct {
	File      string
	Directory string
}

/** Plugin Type */
type WPPlugin struct {
	Name    string
	Path    WPPluginPath
	Version string
	Content string
}

/* WP Plugin Check */
func WPPluginBuildCheck(path string, production bool) {
	fmt.Println("Check Plugin")
	plugin := WPPlugin{}
	plugin.Path.Directory = path

	/** Get Plugin Information */
	PathArray := strings.Split(path, "/")
	plugin.Name = PathArray[len(PathArray)-1]
	FileName := fmt.Sprintf("%s.php", PathArray[len(PathArray)-1])
	PathArray = append(PathArray, FileName)
	plugin.Path.File = strings.Join(PathArray, "/")

	/** Read File */
	content := library.ReadFile(plugin.Path.File)

	/** Read Comment Block */
	regexcommentblock := regexp.MustCompile("(?s)//.*?\n|/\\*.*?\\*/")
	comments := strings.Split(regexcommentblock.FindString(string(content)), "\n")
	for _, s := range comments {
		s = strings.Replace(s, "*", "", -1)
		if strings.Contains(s, "Name:") {
			s = strings.Replace(s, "Plugin Name:", "", -1)
			plugin.Name = strings.Join(strings.Fields(s), " ")
		}
		if strings.Contains(s, "Version:") {
			s = strings.Replace(s, " ", "", -1)
			plugin.Version = strings.Replace(s, "Version:", "", -1)
		}
	}

	/** Check occurance (config.json) */
	FileName = "config.json"
	content = library.ReadFile(plugin.Path.Directory + "/" + FileName)
	res, err := regexp.Match(plugin.Version, content)
	fmt.Println("Plugin Version Match", FileName, ":", res, "with Error :", err)
}
