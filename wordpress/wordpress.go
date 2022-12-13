package wordpress

import (
	"aspri/library"
	"fmt"
	"strings"
)

/** Initiate WordPress Function */
func InitiateWordPressFunction(flags library.Flag) {
	/** Refactor Plugin */
	if *flags.WPRefactor && *flags.Path != "" && *flags.From != "" && *flags.To != "" {
		WPRefactor(*flags.Path, *flags.From, *flags.To)
	}
	/** WP Plugin Build Check */
	if *flags.WPPluginBuildCheck && *flags.Path != "" {
		WPPluginBuildCheck(*flags.Path, *flags.Production)
	}
}

/* Refactor Plugin */
func WPRefactor(path string, fromName string, toName string) {
	fmt.Print("Refactor Plugin: ", fromName, " to ", toName)
	library.SearchandReplaceinDir(path, fromName, toName)
	library.SearchandReplaceinDir(path, strings.ToUpper(fromName), strings.ToUpper(toName))
	library.SearchandReplaceinDir(path, strings.ToLower(fromName), strings.ToLower(toName))
}
