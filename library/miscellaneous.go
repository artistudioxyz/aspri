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
	/** removeFilesExceptExtensions */
	if *flags.File && *flags.Remove && len(*flags.Ext) > 0 {
		RemoveFilesExceptExtensions(*flags.Path, *flags.Ext, *flags.Except)
	}
	/** Delete Directory or Files in Path Matching Filename */
	if *flags.Dir && *flags.Remove && len(*flags.Filename) > 0 {
		DeleteDirectoriesorFilesinPath(true, *flags.Path, *flags.Filename)
	}
	if *flags.File && *flags.Remove && len(*flags.Filename) > 0 {
		DeleteDirectoriesorFilesinPath(false, *flags.Path, *flags.Filename)
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

/** Remove Files Except Specified Extensions */
func RemoveFilesExceptExtensions(root string, allowedExtensions []string, exception []string) error {
	if root == "" {
		CurrentDirectory, _ := os.Getwd()
		root = CurrentDirectory
	}

	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("❌ ", err)
		}
		if !info.IsDir() {
			ext := filepath.Ext(info.Name())
			if !SliceContainsString(allowedExtensions, ext) && !SliceContainsString(exception, info.Name()) {
				err := os.Remove(path)
				if err != nil {
					return err
				}
				fmt.Println("✅ Successfully remove files except extensions", allowedExtensions, "in", info.Name())
			}
		}
		return nil
	})
}

/** Delete Directory or Files in Path Matching Filename */
func DeleteDirectoriesorFilesinPath(isdir bool, root string, filenames []string) error {
	if root == "" {
		CurrentDirectory, _ := os.Getwd()
		root = CurrentDirectory
	}

	// Walk through the directory tree
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// If the path is a directory and it has the correct name, delete it
		if isdir && info.IsDir() && SliceContainsString(filenames, info.Name()) {
			err = os.RemoveAll(path)
			if err != nil {
				return err
			}
			fmt.Println("✅ Successfully remove directories nested by filename", info.Name(), "in", root)
		} else if !isdir && !info.IsDir() && SliceContainsString(filenames, info.Name()) {
			err = os.RemoveAll(path)
			if err != nil {
				return err
			}
			fmt.Println("✅ Successfully remove files nested by filename", info.Name(), "in", root)
		}

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
