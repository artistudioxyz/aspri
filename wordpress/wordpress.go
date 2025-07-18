package wordpress

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/artistudioxyz/aspri/library"
	"os"
	"path/filepath"
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
	if *flags.WPRefactor && *flags.From != "" && *flags.To != "" {
		WPRefactor(*flags.Path, *flags.From, *flags.To, *flags.Type)
	}
	/** WP Clean Project Files for Production */
	if *flags.WPClean && *flags.Type != "" {
		CleanProjectFilesforProduction(*flags.Path, *flags.Type)
	}
	/** WP Plugin Build Check */
	if *flags.WPPluginBuildCheck {
		plugin := GetPluginInformation(*flags.Path)
		CheckProjectVersion(plugin)
	}
	/** WP Theme Build Check */
	if *flags.WPThemeBuildCheck {
		theme := GetThemeInformation(*flags.Path)
		CheckProjectVersion(theme)
	}
	/** WP Plugin Build */
	if *flags.WPPluginBuild && *flags.Type != "" {
		plugin := GetPluginInformation(*flags.Path)
		CheckProjectVersion(plugin)
		CleanProjectFilesforProduction(*flags.Path, *flags.Type)
		SetConfigProduction(*flags.Path, true)
	}
	/** WP Theme Build */
	if *flags.WPThemeBuild && *flags.Type != "" {
		theme := GetThemeInformation(*flags.Path)
		CheckProjectVersion(theme)
		CleanProjectFilesforProduction(*flags.Path, *flags.Type)
		CleanVendorDirandFilesforProduction(*flags.Path, "theme")
		SetConfigProduction(*flags.Path, true)
	}
	/** WP Plugin Release */
	if *flags.WPPluginRelease && *flags.To != "" {
		plugin := GetPluginInformation(*flags.Path)
		fmt.Println("📦 Project Name:", plugin.Name)
		fmt.Println("📦 Project Version:", plugin.Version)
		library.SearchandReplaceDirectory(plugin.Path.File, plugin.Version, *flags.To, 1)
		library.SearchandReplaceDirectory(*flags.Path+"/readme.txt", plugin.Version, *flags.To, 1)
		library.SearchandReplaceDirectory(*flags.Path+"/config.json", plugin.Version, *flags.To, -1)
		library.SearchandReplaceDirectory(*flags.Path+"/package.json", plugin.Version, *flags.To, -1)
	}
	/** WP Tag Trunk */
	if *flags.WPTagTrunk {
		plugin := GetPluginInformation(*flags.Path + "/trunk")
		fmt.Println("📦 Project Name:", plugin.Name)
		fmt.Println("📦 Project Version:", plugin.Version)
		library.ExecCommand("mkdir", "-p", *flags.Path+"/tags/"+plugin.Version)
		library.ExecCommand("rsync", "--delete", "-av", *flags.Path+"/trunk/", *flags.Path+"/tags/"+plugin.Version)
	}
}

/* Refactor Dot Framework */
func WPRefactor(path string, fromName string, toName string, BuildType string) {
	// If build type is not defined, set it to plugin.
	if BuildType == "" {
		BuildType = "plugin"
	}

	// Do refactor.
	var shell bytes.Buffer
	fmt.Println("Refactor Plugin: ", fromName, " to ", toName)
	library.SearchandReplaceDirectory(path, fromName, toName, -1)
	library.SearchandReplaceDirectory(path, strings.ToUpper(fromName), strings.ToUpper(toName), -1)
	library.SearchandReplaceDirectory(path, strings.ToLower(fromName), strings.ToLower(toName), -1)
	if BuildType == "plugin" {
		shell.WriteString(library.GetShellRemoveFunction(path + "/src/Theme.php"))
		library.RenameFile(path+"/dot.php", path+string(filepath.Separator)+strings.ToLower(toName)+".php")
	} else if BuildType == "theme" {
		shell.WriteString("mv " + path + "/dot.php " + path + "/functions.php")
		shell.WriteString(library.GetShellRemoveFunction(path + "/src/Plugin.php"))
		library.SearchandReplaceDirectory(path, fmt.Sprintf("%s_PLUGIN", strings.ToUpper(toName)), fmt.Sprintf("%s_THEME", strings.ToUpper(toName)), -1)
		library.SearchandReplaceDirectory(path, fmt.Sprintf("%s Plugins", strings.ToUpper(toName)), fmt.Sprintf("%s Theme", strings.ToUpper(toName)), -1)
		library.SearchandReplaceDirectory(path, fmt.Sprintf("%s Plugin", strings.ToUpper(toName)), fmt.Sprintf("%s Theme", strings.ToUpper(toName)), -1)
		library.SearchandReplaceDirectory(path, "use Helper\\Model;", "", -1)

		/** Remove Model */
		shell.WriteString(library.GetShellRemoveFunction(path + "/src/WordPress/Model"))
		shell.WriteString(library.GetShellRemoveFunction(path + "/src/WordPress/Helper/Model"))
		shell.WriteString(library.GetShellRemoveFunction(path + "/src/WordPress/Page/MenuPage.php"))
		shell.WriteString(library.GetShellRemoveFunction(path + "/src/WordPress/Page/SubmenuPage.php"))
	}
	cmd := [...]string{"bash", "-c", shell.String()}
	library.ExecCommand(cmd[:]...)
}

/** CleanVendorDirandFilesforProduction */
func CleanVendorDirandFilesforProduction(path string, BuildType string) {
	/** Delete Directories and Files */
	library.RemoveFilesExceptExtensions(path+"/vendor/", []string{".php"}, []string{})
	library.DeleteDirectoriesorFilesinPath(path+"/vendor/",
		[]string{
			"languages",
			"plugins",
			".github",
			".husky",
		},
		[]string{
			"example.php",
			"index.php",
		})

	if BuildType == "theme" {
		library.DeleteDirectoriesorFilesinPath(path+"/vendor/",
			[]string{},
			[]string{
				"Email.php",
				"Model.php",
			})
	}

	fmt.Println("✅ Success Cleanup Vendor Directories and Files for Production")
}

/** CleanProjectFilesforProduction */
func CleanProjectFilesforProduction(path string, buildType string) {
	var remove bytes.Buffer
	var Files = []string{
		// Operating System
		".DS_Store",

		/** Git */
		".gitignore",

		/** Hooks */
		".husky",
		".editorconfig",
		".eslintignore",
		".eslintrc.json",
		".prettierignore",
		".prettierrc",
		".prettierrc.json",
		".release-it.json",
		"commitlint.config.js",

		/** Vendor */
		"node_modules",

		/** Tests */
		"tests",
		"tests-selenium",

		/** Assets */
		"assets/css",
		"assets/js",
		"assets/ts",
		"assets/blocks",
		"assets/components",
		"assets/dist/css/tailwind.css",
		"assets/dist/css/tailwind.min.css",
		"assets/dist/ts",
		"assets/dist/*.map",

		/** Development Files */
		"helper.php",
		"Gruntfile.js",
		"components.json",
		"composer.json",
		"composer.lock",
		"originalassets.js",
		"package.json",
		"package-lock.json",
		"pnpm-lock.yaml",
		"tailwind-default.config.js",
		"tailwind.config.js",
		"tailwindcsssupport.js",
		"tsconfig.json",
		"webpack.config.js",
		"README.md",
		"sniffer.txt",
		"vite.config.js",
	}
	var FilesforGithub = []string{ // Lists of files that is required for GitHub
		"README.md",
	}

	/** Filter & Generate Command */
	for _, f := range Files {
		if buildType == "github" {
			ForGithub := library.SliceContainsString(FilesforGithub, f)
			if !ForGithub {
				remove.WriteString(library.GetShellRemoveFunction(path + string(filepath.Separator) + f))
			}
		} else {
			remove.WriteString(library.GetShellRemoveFunction(path + string(filepath.Separator) + f))
		}
	}
	cmd := [...]string{"bash", "-c", remove.String()}
	library.ExecCommand(cmd[:]...)

	fmt.Println("✅ Success Cleanup Project Files")
}

/** SetConfigProduction */
func SetConfigProduction(path string, production bool) {
	plugin := GetPluginInformation(path)
	FileName := "config.json"

	// Check if file exists
	configPath := plugin.Path.Directory + string(filepath.Separator) + FileName
	if _, err := os.Stat(configPath); err != nil {
		return
	}

	// Get Content
	content := library.ReadFile(configPath)

	// Read and Change Value
	var objmap map[string]interface{}
	if err := json.Unmarshal(content, &objmap); err != nil {
		panic(err)
	}
	objmap["production"] = production
	jsonStr, _ := json.Marshal(objmap)
	library.WriteFile(plugin.Path.Directory+string(filepath.Separator)+FileName, string(jsonStr))

	fmt.Println("✅ Success set production config to", production)
}

// Read comment block
func ReadCommentBlock(project WPProject) WPProject {
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

	return project
}

/** Check Version */
func CheckProjectVersion(project WPProject) {
	// Ouput project name and version
	fmt.Println("📦 Project Name:", project.Name)
	fmt.Println("📦 Project Version:", project.Version)

	/** Check occurrence (readme.txt) */
	FileName := "readme.txt"
	content := library.ReadFile(project.Path.Directory + string(filepath.Separator) + FileName)
	regexversion := regexp.MustCompile(project.Version)
	matches := regexversion.FindAllStringIndex(string(content), 2)
	if len(matches) >= 1 {
		fmt.Println("✅ Plugin Version Match", FileName)
	} else {
		panic("❌ Plugin Version Do Not Match " + FileName)
	}

	// Check occurrence (config.json)
	FileName = "config.json"
	configPath := project.Path.Directory + string(filepath.Separator) + FileName
	if _, err := os.Stat(configPath); err == nil {
		content = library.ReadFile(configPath)
		if res, err := regexp.Match(project.Version, content); res {
			fmt.Println("✅ Plugin Version Match", FileName)
		} else if err != nil {
			panic(err)
		} else {
			fmt.Println("❌ Plugin Version Do Not Match " + FileName)
			panic("Version mismatch in " + FileName)
		}
	}

	// Check occurrence (package.json)
	FileName = "package.json"
	if _, err := os.Stat(project.Path.Directory + string(filepath.Separator) + FileName); err == nil {
		content = library.ReadFile(project.Path.Directory + string(filepath.Separator) + FileName)
		res, err := regexp.Match(project.Version, content)
		if res {
			fmt.Println("✅ Plugin Version Match", FileName)
		} else {
			fmt.Println("❌ Plugin Version Do Not Match " + FileName)
			panic(err)
		}
	}
}
