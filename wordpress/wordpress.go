package wordpress

import (
	"aspri/library"
	"bytes"
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
	var remove bytes.Buffer
	var Files = []string{
		/** Git */
		".git",
		".gitignore",

		/** Vendor */
		"node_modules",

		/** Tests */
		"tests-selenium",

		/** Assets */
		"assets/css",
		"assets/js",
		"assets/ts",
		"assets/components",
		"assets/build/css/tailwind.min.css",
		"assets/build/ts",
		"assets/build/*.map",

		/** Development Files */
		"livereload.php",
		"Gruntfile.js",
		"composer.json",
		"composer.lock",
		"package-lock.json",
		"package.json",
		"tailwind-default.config.js",
		"tailwind.config.js",
		"tailwindcsssupport.js",
		"tsconfig.json",
		"webpack.config.js",
		"CHANGELOG.md",
		"DOCS.md",
		"README.md",
	}
	var FilesforGithub = []string{ // Lists of files that is required for GitHub
		".gitignore",
		"README.md",
	}

	/** Filter & Generate Command */
	for _, f := range Files {
		if buildType == "github" {
			ForGithub := library.SliceContainsString(FilesforGithub, f)
			if !ForGithub {
				remove.WriteString(library.GetShellRemoveFunction(path + "/" + f))
			}
		} else {
			remove.WriteString(library.GetShellRemoveFunction(path + "/" + f))
		}
	}
	cmd := [...]string{"bash", "-c", remove.String()}
	library.ExecCommand(cmd[:]...)
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
