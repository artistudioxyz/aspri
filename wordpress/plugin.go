package wordpress

import (
	"flag"
	"fmt"
)

var (
	/** WordPress */
	WPPluginBuildCheckFlag = flag.Bool("wp-plugin-build-check", false, "WP Check Plugin Comply with Directory")
)

/* WP Plugin Check */
func WPPluginBuildCheck(path string) {
	fmt.Println("Check Plugin")

	fmt.Println(path)
	path = "mantap"
	fmt.Println(path)
}
