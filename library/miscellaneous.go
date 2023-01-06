package library

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

/** Initiate Miscellaneous Function */
func InitiateMiscellaneousFunction(flags Flag) {
	/** removeFilesExceptExtensions */
	if *flags.File && *flags.Remove && len(*flags.Ext) > 0 {
		RemoveFilesExceptExtensions(*flags.Path, *flags.Ext)
	}
	/** Delete Directory by Regex */
	if *flags.Dir && *flags.Remove && *flags.Regex != "" {
		err := DeleteDirectoryByRegex(*flags.Path, *flags.Regex)
		if err != nil {
			fmt.Println("❌ ", err)
		}
	}
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

/** Delete Directory by Regex */
func DeleteDirectoryByRegex(path string, regexString string) error {
	if path == "" {
		CurrentDirectory, _ := os.Getwd()
		path = CurrentDirectory
	}

	regex, err := regexp.Compile(regexString)
	if err != nil {
		panic(err)
	}

	// Walk the directory tree and remove all directories that match the pattern
	return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && regex.MatchString(info.Name()) {
			return os.RemoveAll(path)
		}

		fmt.Println("✅ Successfully remove directory by regex", path, "in", regexString)
		return nil
	})
}

/** Remove Files Except Specified Extensions */
func RemoveFilesExceptExtensions(root string, allowedExtensions []string) error {
	if root == "" {
		CurrentDirectory, _ := os.Getwd()
		root = CurrentDirectory
	}

	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			ext := filepath.Ext(info.Name())
			if !SliceContainsString(allowedExtensions, ext) {
				err := os.Remove(path)
				if err != nil {
					return err
				}
			}
		}
		fmt.Println("✅ Successfully remove files except extensions", root, "in", allowedExtensions)
		return nil
	})
}

/** Search and Replace in Directory or File */
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
