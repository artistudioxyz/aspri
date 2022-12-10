package library

import (
	"flag"
	"fmt"
	"strings"
)

var (
	/** WordPress */
	buildWPPluginFlag = flag.Bool("wp-plugin-build", false, "Build WP Plugin")
)

/** Initiate WordPress Function */
func InitiateWordPressFunction() {
	/** Build WP Plugin */
	if *buildWPPluginFlag && *RefactorPath != "" && *RefactorFromFlag != "" && *RefactorToFlag != "" {
		ProjectRefactorPlugin(*RefactorPath, *RefactorFromFlag, *RefactorToFlag)
	}
}

/* Refactor Plugin */
func ProjectRefactorPlugin(path string, fromName string, toName string) {
	fmt.Print("Refactor Plugin: ", fromName, " to ", toName)
	SearchandReplaceinDir(path, fromName, toName)
	SearchandReplaceinDir(path, strings.ToUpper(fromName), strings.ToUpper(toName))
	SearchandReplaceinDir(path, strings.ToLower(fromName), strings.ToLower(toName))
}
