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

/** Initiate Projects Function */
func InitiateProjectFunction() {
	/** Build WP Plugin */
	if *buildWPPluginFlag && *RefactorPath != "" && *RefactorFromFlag != "" && *RefactorToFlag != "" {
		RefactorPlugin(*RefactorPath, *RefactorFromFlag, *RefactorToFlag)
	}
}

/* Refactor Plugin */
func RefactorPlugin(path string, fromName string, toName string) {
	fmt.Print("Refactor Plugin: ", fromName, " to ", toName)
	SearchandReplaceinDir(path, fromName, toName)
	SearchandReplaceinDir(path, strings.ToUpper(fromName), strings.ToUpper(toName))
	SearchandReplaceinDir(path, strings.ToLower(fromName), strings.ToLower(toName))
}
