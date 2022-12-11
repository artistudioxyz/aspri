package main

import (
	"aspri/library"
	"aspri/wordpress"
	"flag"
)

var (
	Version = "1.0.0"
)

func main() {
	/** Grab The Flag */
	flags := library.GetFlag()
	flag.Parse()

	/** Information Flag */
	handleVersionFlag(*flags.Version)
	handleHelpFlag()

	///** Initiate Library */
	library.InitiateDockerFunction(flags)
	library.InitiateGitFunction(flags)
	library.InitiateMiscellaneousFunction(flags)
	library.InitiateQuoteFunction(flags)

	///** Initiate WordPress */
	wordpress.InitiateWordPressFunction(flags)
}
