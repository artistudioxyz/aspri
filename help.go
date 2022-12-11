package main

import (
	"flag"
	"fmt"
)

var (
	helpFlag = flag.Bool("help", false, "show usage message")
)

/** Documentation Help */
const HelpText = `
(｡◕‿‿◕｡) ASPRI (Asisten Pribadi)
Collection of scripts and library to speed up sotware development process 
Learn More: https://github.com/artistudioxyz/aspri
`

/** Help Flag */
func handleHelpFlag() {
	if *helpFlag {
		fmt.Println(HelpText)
		flag.Usage()
		return
	}
}
