package main

import (
	"github.com/artistudioxyz/aspri/library"
	"github.com/artistudioxyz/aspri/wordpress"
)

func main() {
	// Grab the Flag
	flags := library.GetFlag()

	// Initiate Library
	library.InitiateChatGPTFunction(flags)
	library.InitiateContribution(flags)
	library.InitiateDockerFunction(flags)
	library.InitiateFileFunction(flags)
	library.InitiateDirectoryFunction(flags)
	library.InitiateGitFunction(flags)
	library.InitiateHelpFunction(flags)
	library.InitiateMarkdownFunction(flags)
	library.InitiateMiscellaneousFunction(flags)
	library.InitiateNoIPFunction(flags)
	library.InitiateJsonFunction(flags)
	library.InitiatePHPFunction(flags)
	library.InitiatePHPCSFunction(flags)
	library.InitiateRsyncFunction(flags)
	library.InitiateSyncthingFunction(flags)
	library.InitiateQuoteFunction(flags)
	library.InitiateYouTubeFunction(flags)

	// Initiate WordPress
	wordpress.InitiateWordPressFunction(flags)
}
