package library

import (
	"log"
	"os/exec"
)

/** Documentation Help --help */
const HelpText = `
Usage:

 gorename (-from <spec> | -offset <file>:#<byte-offset>) -to <name> [-force]

You must specify the object (named entity) to rename using the -offset
or -from flag.  Exactly one must be specified.

Flags:

-offset    specifies the filename and byte offset of an identifier to rename.
           This form is intended for use by text editors.

`

/* Run custom bin command */
func ExecCommand(args ...string) string {
	cmd := exec.Command(args[0], args[1:]...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Runing failed: %v", err)
	}
	return string(b[:])
}
