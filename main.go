package main

import (
	"flag"
	"github.com/artistudioxyz/aspri/library"
	"github.com/artistudioxyz/aspri/wordpress"
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
