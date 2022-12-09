package main

import (
	"aspri/library"
	"flag"
	"fmt"
)

var (
	/** WordPress */
	buildWPPluginFlag = flag.Bool("build-wp-plugin", false, "Build WP Plugin")

	/** Function - Refactor */
	RefactorPath     = flag.String("path", "", "working directory path")
	RefactorFromFlag = flag.String("from", "", "identifier to be renamed; see -help for formats")
	RefactorToFlag   = flag.String("to", "", "new name for identifier")

	/** Help */
	helpFlag = flag.Bool("help", false, "show usage message")
)

func main() {
	flag.Parse()
	fmt.Print(`(｡◕‿‿◕｡) ASPRI (Asisten Pribadi)
Collection of scripts and library to speed up sotware development process 
Learn More: https://github.com/artistudioxyz/aspri
`)

	/** Help */
	if *helpFlag {
		fmt.Print(library.HelpText)
		return
	}

	/** Build for WP Plugin */
	if *buildWPPluginFlag && *RefactorPath != "" && *RefactorFromFlag != "" && *RefactorToFlag != "" {
		library.RefactorPlugin(*RefactorPath, *RefactorFromFlag, *RefactorToFlag)
	}
}
