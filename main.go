package main

import (
	"aspri/library"
	"aspri/wordpress"
	"flag"
)

func main() {
	/** Grab The Flag */
	flags := library.GetFlag()
	flag.Parse()

	/** Information Flag */
	handleHelpFlag()

	///** Initiate Library */
	library.InitiateDockerFunction(flags)
	library.InitiateGitFunction(flags)
	library.InitiateMiscellaneousFunction(flags)
	library.InitiateQuoteFunction(flags)

	///** Initiate WordPress */
	wordpress.InitiateWordPressFunction(flags)
}
