package library

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

/** Initiate Miscellaneous Function */
func InitiateMiscellaneousFunction(flags Flag) {
	/** Search and Replace */
	if *flags.SearchandReplace && *flags.Path != "" && *flags.From != "" && *flags.To != "" {
		SearchandReplace(*flags.Path, *flags.From, *flags.To)
	}
	/** Self Update */
	if *flags.SelfUpdate {
		cmd := [...]string{"bash", "-c", "go get github.com/artistudioxyz/aspri"}
		ExecCommand(cmd[:]...)
	}
}

/**
* Search and Replace in Directory or File
* - Equivalent to : `LC_ALL=C find {path} -type f -name '*' -exec sed -i '' s/{from}/{to}/g {} +;`
* - Equivalent to : `sed -i 's#{from}#{to}#g' {path}`
 */
func SearchandReplace(path string, from string, to string) {
	/** Define Delimiter */
	DM := ""
	if strings.Contains(from, "/") || strings.Contains(from, "/") {
		DM = "#"
	} else {
		DM = "/"
	}

	/** Generate Shell Command based on Path Type (Directory or File) */
	find := fmt.Sprintf("s%s%s%s%s%sg", DM, from, DM, to, DM)
	fileInfo, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	if fileInfo.IsDir() {
		os := runtime.GOOS
		switch os {
		case "darwin": /** MAC operating system */
			find = fmt.Sprintf(`LC_ALL=C find %s -type f -name '*' -exec sed -i '' "%s" {} +;`, path, find)
		case "linux": /** Linux */
			find = fmt.Sprintf(`LC_ALL=C find %s -type f -name '*' -exec sed -i "%s" {} +;`, path, find)
		default:
			find = fmt.Sprintf(`LC_ALL=C find %s -type f -name '*' -exec sed -i "%s" {} +;`, path, find)
		}
	} else {
		find = fmt.Sprintf("sed -i '' -e '%s' %s", find, path)
	}
	cmd := [...]string{"bash", "-c", find}
	ExecCommand(cmd[:]...)

	fmt.Println("✅ Success Search and Replace", from, "to", to, "in", path)
}
