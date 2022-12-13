package wordpress

import (
	"aspri/library"
	"encoding/json"
	"fmt"
	"strings"
)

/** Initiate WordPress Function */
func InitiateWordPressFunction(flags library.Flag) {
	/** Refactor Plugin */
	if *flags.WPRefactor && *flags.Path != "" && *flags.From != "" && *flags.To != "" {
		WPRefactor(*flags.Path, *flags.From, *flags.To)
	}
	/** WP Plugin Build */
	if *flags.WPPluginBuild && *flags.Path != "" && *flags.Type != "" {
		CleanProjectFilesforProduction(*flags.Path, *flags.Type)
		//SetConfigProduction(*flags.Path, true)
	}
	/** WP Plugin Build Check */
	if *flags.WPPluginBuildCheck {
		WPPluginBuildCheck(*flags.Path)
	}
}

/* Refactor Plugin */
func WPRefactor(path string, fromName string, toName string) {
	fmt.Print("Refactor Plugin: ", fromName, " to ", toName)
	library.SearchandReplaceinDir(path, fromName, toName)
	library.SearchandReplaceinDir(path, strings.ToUpper(fromName), strings.ToUpper(toName))
	library.SearchandReplaceinDir(path, strings.ToLower(fromName), strings.ToLower(toName))
}

/** CleanProjectFilesforProduction */
func CleanProjectFilesforProduction(path string, buildType string) {
	var Files = []string{
		".git",
		".gitignore",
	}

	fmt.Println(Files)

	//var buffer bytes.Buffer
	/** Git */
	//buffer.WriteString(fmt.Sprintf("rm -rf %s/.git;", path))
	//if buildType != "github" {
	//	buffer.WriteString(fmt.Sprintf("rm -rf %s/.git;", path))
	//}
	//
	///** Vendor */
	//buffer.WriteString(fmt.Sprintf("rm -rf %s/.git;", path))
	//
	//cmd := [...]string{"bash", "-c", buffer.String()}
	//library.ExecCommand(cmd[:]...)
}

/** SetConfigProduction */
func SetConfigProduction(path string, production bool) {
	plugin := GetPluginInformation(path)
	FileName := "config.json"
	content := library.ReadFile(plugin.Path.Directory + "/" + FileName)

	/** Read and Change Value */
	var objmap map[string]interface{}
	if err := json.Unmarshal(content, &objmap); err != nil {
		panic(err)
	}
	objmap["production"] = production
	jsonStr, _ := json.Marshal(objmap)
	library.WriteFile(plugin.Path.Directory+"/"+FileName, string(jsonStr))

	fmt.Println("✅ Success set production config to", production)
}
