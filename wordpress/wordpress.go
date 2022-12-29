package wordpress

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/artistudioxyz/aspri/library"
	"os"
	"regexp"
	"strings"
)

/** Path Type */
type WPPath struct {
	File      string
	Directory string
}

/** Plugin Type */
type WPProject struct {
	Name    string
	Path    WPPath
	Version string
	Content string
}

/** Initiate WordPress Function */
func InitiateWordPressFunction(flags library.Flag) {
	/** Refactor Plugin */
	if *flags.WPRefactor && *flags.Path != "" && *flags.From != "" && *flags.To != "" && *flags.Type != "" {
		WPRefactor(*flags.Path, *flags.From, *flags.To, *flags.Type)
	}
	/** WP Plugin Build Check */
	if *flags.WPPluginBuildCheck {
		WPPluginBuildCheck(*flags.Path)
	}
	/** WP Theme Build Check */
	if *flags.WPThemeBuildCheck {
		WPThemeBuildCheck(*flags.Path)
	}
	/** WP Plugin Build */
	if *flags.WPPluginBuild && *flags.Type != "" {
		WPPluginBuildCheck(*flags.Path)
		CleanProjectFilesforProduction(*flags.Path, *flags.Type)
		SetConfigProduction(*flags.Path, true)
	}
	/** WP Theme Build */
	if *flags.WPThemeBuild && *flags.Type != "" {
		WPThemeBuildCheck(*flags.Path)
		CleanProjectFilesforProduction(*flags.Path, *flags.Type)
		SetConfigProduction(*flags.Path, true)
	}
}

/* Refactor Dot Framework */
func WPRefactor(path string, fromName string, toName string, BuildType string) {
	var remove bytes.Buffer
	fmt.Print("Refactor Plugin: ", fromName, " to ", toName)
	library.SearchandReplace(path, fromName, toName)
	library.SearchandReplace(path, strings.ToUpper(fromName), strings.ToUpper(toName))
	library.SearchandReplace(path, strings.ToLower(fromName), strings.ToLower(toName))
	if BuildType == "plugin" {
		remove.WriteString(library.GetShellRemoveFunction(path + "/src/Framework/Theme.php"))
	} else if BuildType == "theme" {
		remove.WriteString(library.GetShellRemoveFunction(path + "/src/Framework/Plugin.php"))
		library.SearchandReplace(path, fmt.Sprintf("%s_PLUGIN", strings.ToUpper(toName)), fmt.Sprintf("%s_THEME", strings.ToUpper(toName)))
		library.SearchandReplace(path, fmt.Sprintf("%s Plugins", strings.ToUpper(toName)), fmt.Sprintf("%s Theme", strings.ToUpper(toName)))
		library.SearchandReplace(path, fmt.Sprintf("%s Plugin", strings.ToUpper(toName)), fmt.Sprintf("%s Theme", strings.ToUpper(toName)))

		/** Remove Model */
		remove.WriteString(library.GetShellRemoveFunction(path + "/src/WordPress/Model"))
		remove.WriteString(library.GetShellRemoveFunction(path + "/src/WordPress/Helper/Model"))
		remove.WriteString(library.GetShellRemoveFunction(path + "/src/WordPress/Page/MenuPage.php"))
		remove.WriteString(library.GetShellRemoveFunction(path + "/src/WordPress/Page/SubmenuPage.php"))
		//library.SearchandReplace(path, "use Helper\\Page;", "")
	}
	cmd := [...]string{"bash", "-c", remove.String()}
	library.ExecCommand(cmd[:]...)
}

/** CleanProjectFilesforProduction */
func CleanProjectFilesforProduction(path string, buildType string) {
	if path == "" {
		CurrentDirectory, _ := os.Getwd()
		path = CurrentDirectory
	}

	var remove bytes.Buffer
	var Files = []string{
		/** Git */
		".gitignore",

		/** Hooks */
		".husky",
		".editorconfig",
		".eslintignore",
		".eslintrc.json",
		".prettierignore",
		".prettierrc.json",
		".release-it.json",
		"commitlint.config.js",

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
		"originalassets.js",
		"package-lock.json",
		"package.json",
		"tailwind-default.config.js",
		"tailwind.config.js",
		"tailwindcsssupport.js",
		"tsconfig.json",
		"webpack.config.js",
		"DOCS.md",
		"README.md",
	}
	var FilesforGithub = []string{ // Lists of files that is required for GitHub
		".git",
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

	/** Exclude File From .gitignore for BuildType (GitHub) */
	if buildType == "github" {
		library.SearchandReplace(path+"/.gitignore", "/assets/build/", "")
		library.SearchandReplace(path+"/.gitignore", "/assets/vendor/**/*.js", "")
		library.SearchandReplace(path+"/.gitignore", "/assets/vendor/**/*.css", "")
		library.SearchandReplace(path+"/.gitignore", "!/assets/vendor/**/*.min.*", "")
		library.SearchandReplace(path+"/.gitignore", "/vendor/", "")
	}

	fmt.Println("✅ Success Cleanup Project Files")
}

/** SetConfigProduction */
func SetConfigProduction(path string, production bool) {
	if path == "" {
		CurrentDirectory, _ := os.Getwd()
		path = CurrentDirectory
	}

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

/** Check Version */
func CheckProjectVersion(project WPProject) {
	/** Read Comment Block */
	content := library.ReadFile(project.Path.File)
	regexcommentblock := regexp.MustCompile("(?s)//.*?\n|/\\*.*?\\*/")
	comments := strings.Split(regexcommentblock.FindString(string(content)), "\n")
	for _, s := range comments {
		s = strings.Replace(s, "*", "", -1)
		if strings.Contains(s, "Name:") {
			s = strings.Replace(s, "Plugin Name:", "", -1)
			project.Name = strings.Join(strings.Fields(s), " ")
		}
		if strings.Contains(s, "Version:") {
			s = strings.Replace(s, " ", "", -1)
			project.Version = strings.Replace(s, "Version:", "", -1)
		}
	}

	/** Check occurrence (readme.txt) */
	FileName := "readme.txt"
	content = library.ReadFile(project.Path.Directory + "/" + FileName)
	regexversion := regexp.MustCompile(project.Version)
	matches := regexversion.FindAllStringIndex(string(content), 2)
	if len(matches) >= 1 {
		fmt.Println("✅ Plugin Version Match", FileName)
	} else {
		panic("❌ Plugin Version Do Not Match " + FileName)
	}

	/** Check occurrence (config.json) */
	FileName = "config.json"
	if _, err := os.Stat(project.Path.Directory + "/" + FileName); err == nil {
		content = library.ReadFile(project.Path.Directory + "/" + FileName)
		res, err := regexp.Match(project.Version, content)
		if res {
			fmt.Println("✅ Plugin Version Match", FileName)
		} else {
			fmt.Println("❌ Plugin Version Do Not Match " + FileName)
			panic(err)
		}
	}
}
