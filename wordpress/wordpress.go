package wordpress

import (
	"aspri/library"
	"fmt"
	"strings"
)

/** Initiate WordPress Function */
func InitiateWordPressFunction(flags library.Flag) {
	/** Refactor Plugin */
	if *flags.WPRefactor && *library.RefactorPath != "" && *library.RefactorFromFlag != "" && *library.RefactorToFlag != "" {
		WPRefactor(*library.RefactorPath, *library.RefactorFromFlag, *library.RefactorToFlag)
	}
	/** WP Plugin Build Check */
	if *WPPluginBuildCheckFlag {
		//WPPluginBuildCheck(*WPProjectPathFlag)
	}
}

/* Refactor Plugin */
func WPRefactor(path string, fromName string, toName string) {
	fmt.Print("Refactor Plugin: ", fromName, " to ", toName)
	library.SearchandReplaceinDir(path, fromName, toName)
	library.SearchandReplaceinDir(path, strings.ToUpper(fromName), strings.ToUpper(toName))
	library.SearchandReplaceinDir(path, strings.ToLower(fromName), strings.ToLower(toName))
}
