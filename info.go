package main

import "fmt"

/** Handle Info Flag */
func handleVersionFlag(VersionFlag bool) {
	if VersionFlag {
		fmt.Print(Version)
		return
	}
}
