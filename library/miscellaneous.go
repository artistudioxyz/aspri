package library

import (
	"flag"
	"fmt"
)

var (
	/** Refactor */
	RefactorPath     = flag.String("path", "", "working directory path")
	RefactorFromFlag = flag.String("from", "", "string to be replaced")
	RefactorToFlag   = flag.String("to", "", "string replacement")

	/** Search and Replace */
	SearchandReplaceFlag = flag.Bool("search-replace-directory", false, "do search and replace")
)

/** Initiate Miscellaneous Function */
func InitiateMiscellaneousFunction() {
	/** Build WP Plugin */
	if *SearchandReplaceFlag && *RefactorPath != "" && *RefactorFromFlag != "" && *RefactorToFlag != "" {
		SearchandReplaceinDir(*RefactorPath, *RefactorFromFlag, *RefactorToFlag)
	}
}

/** Search and Replace in Directory */
func SearchandReplaceinDir(path string, from string, to string) {
	find := fmt.Sprintf("find %s -type f -name '*' -exec sed -i '' s/%s/%s/g {} +;", path, from, to)
	cmd := [...]string{"bash", "-c", find}
	ExecCommand(cmd[:]...)
}
