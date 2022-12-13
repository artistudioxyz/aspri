package library

import (
	"fmt"
)

/** Initiate Miscellaneous Function */
func InitiateMiscellaneousFunction(flags Flag) {
	/** Build WP Plugin */
	if *flags.SearchandReplaceDirectory && *flags.Path != "" && *flags.From != "" && *flags.To != "" {
		SearchandReplaceinDir(*flags.Path, *flags.From, *flags.To)
	}
}

/**
* Search and Replace in Directory
* - Equivalent to : `find {path} -type f -name '*' -exec sed -i '' s/{from}/{to}/g {} +;`
 */
func SearchandReplaceinDir(path string, from string, to string) {
	find := fmt.Sprintf("find %s -type f -name '*' -exec sed -i '' s/%s/%s/g {} +;", path, from, to)
	cmd := [...]string{"bash", "-c", find}
	ExecCommand(cmd[:]...)
}
