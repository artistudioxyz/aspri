package main

import (
	"aspri/library"
	"aspri/wordpress"
	"flag"
)

var (
	Version = "1.0.0"

	/** Flags */

)

func main() {
	/** Grab The Flag */
	flags := library.GetFlag()
	flag.Parse()

	/** Information Flag */
	handleInfoFlag()
	handleHelpFlag()

	///** Initiate Library */
	library.InitiateDockerFunction(flags)
	library.InitiateGitFunction(flags)
	library.InitiateMiscellaneousFunction(flags)
	//library.InitiateQuoteFunction()

	///** Initiate WordPress */
	wordpress.InitiateWordPressFunction(flags)
}
