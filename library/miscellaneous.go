package library

import (
	"fmt"
	"os"
	"strings"
)

/** Initiate Miscellaneous Function */
func InitiateMiscellaneousFunction(flags Flag) {
	/** Build WP Plugin */
	if *flags.SearchandReplace && *flags.Path != "" && *flags.From != "" && *flags.To != "" {
		SearchandReplace(*flags.Path, *flags.From, *flags.To)
	}
}

/**
* Search and Replace in Directory or File
* - Equivalent to : `find {path} -type f -name '*' -exec sed -i '' s/{from}/{to}/g {} +;`
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
		find = fmt.Sprintf("find %s -type f -name '*' -exec sed -i '' %s {} +;", path, find)
	} else {
		find = fmt.Sprintf("sed -i '%s' %s", find, path)
	}
	cmd := [...]string{"bash", "-c", find}
	ExecCommand(cmd[:]...)

	fmt.Println("✅ Success Search and Replace", from, "to", to, "in", path)
}
