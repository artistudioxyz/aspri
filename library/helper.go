package library

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

/* Run custom bin command */
func ExecCommand(args ...string) string {
	cmd := exec.Command(args[0], args[1:]...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Runing failed: %v", err)
	}
	return string(b)
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
