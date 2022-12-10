package main

import (
	"flag"
	"fmt"
)

var (
	VersionFlag = flag.Bool("version", false, "show current version")
)

/** Handle Info Flag */
func handleInfoFlag() {
	/** Version */
	if *VersionFlag {
		fmt.Print(Version)
		return
	}
}
