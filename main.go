package main

import (
	"aspri/library"
	"flag"
)

var (
	Version = "1.0.0"
)

func main() {
	flag.Parse()

	/** Information Flag */
	handleInfoFlag()
	handleHelpFlag()

	/** Initiate Library */
	library.InitiateDockerFunction()
	library.InitiateGitFunction()
	library.InitiateMiscellaneousFunction()
	library.InitiateQuoteFunction()
	library.InitiateWordPressFunction()
}
