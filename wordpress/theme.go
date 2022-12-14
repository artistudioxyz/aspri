package wordpress

import (
	"aspri/library"
	"fmt"
	"os"
	"regexp"
	"strings"
)

/* WP Theme Check */
func WPThemeBuildCheck(path string) {
	fmt.Println("🔍 Check Theme")
	if path == "" {
		CurrentDirectory, _ := os.Getwd()
		path = CurrentDirectory
	}
	plugin := GetPluginInformation(path)

	/** Read Comment Block */
	content := library.ReadFile(plugin.Path.File)
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

	/** Check occurrence (readme.txt) */
	FileName := "readme.txt"
	content = library.ReadFile(plugin.Path.Directory + "/" + FileName)
	regexversion := regexp.MustCompile(plugin.Version)
	matches := regexversion.FindAllStringIndex(string(content), 2)
	if len(matches) == 2 {
		fmt.Println("✅ Plugin Version Match", FileName)
	} else {
		panic("❌ Plugin Version Do Not Match " + FileName)
	}

	/** Check occurrence (config.json) */
	FileName = "config.json"
	if _, err := os.Stat(plugin.Path.Directory + "/" + FileName); err == nil {
		content = library.ReadFile(plugin.Path.Directory + "/" + FileName)
		res, err := regexp.Match(plugin.Version, content)
		if res {
			fmt.Println("✅ Plugin Version Match", FileName)
		} else {
			fmt.Println("❌ Plugin Version Do Not Match " + FileName)
			panic(err)
		}
	}
}
