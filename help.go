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

Usage:

[Docker](library/docker.go) :
- Stop and Remove Container : '--docker -prune -id {identifier}' 
- Compose restart (down & up) : '--docker-compose-restart'

[Git](library/git.go) :
- Commit and Push : '--git -push -m {message}'

[Quotes](library/quotes.go) :
- Quote of the day : '--quote-of-the-day'

[Miscellaneous](library/miscellaneous.go) :
- Search and Replace in Directory : '--search-replace-directory --path {workdir} -from {text} -to {text}'

[WordPress](library/wordpress.go) :
- Build WP Plugin : '--wp-plugin-build --path {workdir} -from {namespace} -to {namespace}'

`

/** Help Flag */
func handleHelpFlag() {
	if *helpFlag {
		fmt.Print(HelpText)
		return
	}
}
