package library

import (
	"fmt"
	flag "github.com/spf13/pflag"
)

/** Documentation Help */
const HelpText = `(｡◕‿‿◕｡) ASPRI (Asisten Pribadi)
Collection of scripts and library to speed up sotware development process
Learn More: https://github.com/artistudioxyz/aspri
`

/** Help Flag */
func InitiateHelpFunction(flags Flag) {
	if *flags.Help {
		fmt.Println(HelpText)
		flag.Usage()
		return
	}
}
