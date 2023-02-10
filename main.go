package main

import (
	"github.com/artistudioxyz/aspri/library"
	"github.com/artistudioxyz/aspri/wordpress"
)

func main() {
	/** Grab The Flag */
	flags := library.GetFlag()

	/** Information Flag */
	handleHelpFlag()

	///** Initiate Library */
	library.InitiateChatGPTFunction(flags)
	library.InitiateDockerFunction(flags)
	library.InitiateGitFunction(flags)
	library.InitiateMarkdownFunction(flags)
	library.InitiateMiscellaneousFunction(flags)
	library.InitiateNoIPFunction(flags)
	library.InitiatePHPFunction(flags)
	library.InitiatePHPCSFunction(flags)
	library.InitiateSyncthingFunction(flags)
	library.InitiateQuoteFunction(flags)

	///** Initiate WordPress */
	wordpress.InitiateWordPressFunction(flags)
}
