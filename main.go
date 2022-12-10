package main

import (
	"aspri/library"
	"flag"
	"fmt"
)

var (
	/** Extras */
	helpFlag = flag.Bool("help", false, "show usage message")
)

func main() {
	flag.Parse()

	/** Help */
	if *helpFlag {
		fmt.Print(library.HelpText)
		return
	}

	/** Initiate Library */
	library.InitiateDockerFunction()
	library.InitiateGitFunction()
	library.InitiateMiscellaneousFunction()
	library.InitiateQuoteFunction()
	library.InitiateWordPressFunction()
}
