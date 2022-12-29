package library

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

/** Initiate Miscellaneous Function */
func InitiateMiscellaneousFunction(flags Flag) {
	/** Search and Replace */
	if *flags.SearchandReplace && *flags.From != "" && *flags.To != "" {
		SearchandReplace(*flags.Path, *flags.From, *flags.To)
	}
	/** Self Update */
	if *flags.SelfUpdate {
		fmt.Println("✅ Doing self update")
		cmd := [...]string{"bash", "-c", "go get github.com/artistudioxyz/aspri"}
		ExecCommand(cmd[:]...)
	}
}

/**
* Search and Replace in Directory or File
 */
func SearchandReplace(path string, from string, to string) {
	if path == "" {
		CurrentDirectory, _ := os.Getwd()
		path = CurrentDirectory
	}

	/** Search and Replace */
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			newData := strings.Replace(string(data), from, to, -1)
			err = ioutil.WriteFile(path, []byte(newData), 0644)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("❌ ", err)
	} else {
		fmt.Println("✅ Success Search and Replace", from, "to", to, "in", path)
	}
}
