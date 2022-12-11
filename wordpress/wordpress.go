package wordpress

import (
	"aspri/library"
	"flag"
	"fmt"
	"strings"
)

var (
	/** WordPress */
	WPRefactorFlag = flag.Bool("wp-refactor", false, "Refactor Library")
)

/** Initiate WordPress Function */
func InitiateWordPressFunction() {
	/** Refactor Plugin */
	if *WPRefactorFlag && *library.RefactorPath != "" && *library.RefactorFromFlag != "" && *library.RefactorToFlag != "" {
		WPRefactor(*library.RefactorPath, *library.RefactorFromFlag, *library.RefactorToFlag)
	}
	/** WP Plugin Build Check */
	if *WPPluginBuildCheckFlag {
		WPPluginBuildCheck()
	}
}

/* Refactor Plugin */
func WPRefactor(path string, fromName string, toName string) {
	fmt.Print("Refactor Plugin: ", fromName, " to ", toName)
	library.SearchandReplaceinDir(path, fromName, toName)
	library.SearchandReplaceinDir(path, strings.ToUpper(fromName), strings.ToUpper(toName))
	library.SearchandReplaceinDir(path, strings.ToLower(fromName), strings.ToLower(toName))
}
