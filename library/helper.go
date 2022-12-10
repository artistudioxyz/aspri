package library

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

/** Documentation Help --help */
const HelpText = `
(｡◕‿‿◕｡) ASPRI (Asisten Pribadi)
Collection of scripts and library to speed up sotware development process 
Learn More: https://github.com/artistudioxyz/aspri

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

/** Call an API endpoint with Method GET */
func getDataFromAPI(url string) []byte {
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}
