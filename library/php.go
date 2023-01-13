package library

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

/** PHP Class */
type PHPClass struct {
	Name string
	Path string
}

/** PHP Function */
type PHPFunction struct {
	Name string
	Path string
}

/** Initiate PHP Function */
func InitiatePHPFunction(flags Flag) {
	/** List PHP Classes */
	if *flags.PHP && *flags.ListClass {
		classes, _ := ListPHPClasses(*flags.Path)
		for _, class := range classes {
			fmt.Printf("📟 Class Name %s in (%s)\n", class.Name, class.Path)
		}
	}
	/** List PHP Function */
	if *flags.PHP && *flags.ListFunction {
		functions, _ := ListPHPFunctions(*flags.Path)
		for _, function := range functions {
			fmt.Printf("📟 Function Name %s (%s)\n", function.Name, function.Path)
		}
	}
}

/** Function to List PHP Classes inside Directory and Subdirectory */
func ListPHPClasses(root string) ([]PHPClass, error) {
	if root == "" {
		CurrentDirectory, _ := os.Getwd()
		root = CurrentDirectory
	}

	var classes []PHPClass
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".php") {
			// Open the PHP file
			f, err := os.Open(path)
			if err != nil {
				fmt.Println("❌ ", err)
				return err
			}
			defer f.Close()
			// Use a regular expression to find class definitions
			scanner := bufio.NewScanner(f)
			re := regexp.MustCompile(`class\s+([a-zA-Z0-9_]+)`)
			for scanner.Scan() {
				line := scanner.Text()
				matches := re.FindStringSubmatch(line)
				if len(matches) > 1 {
					class := PHPClass{
						Name: matches[1],
						Path: path,
					}
					classes = append(classes, class)
				}
			}
			if err := scanner.Err(); err != nil {
				fmt.Println("❌ ", err)
				return err
			}
		}
		return nil
	})
	return classes, err
}

/** List PHP Functions */
func ListPHPFunctions(root string) ([]PHPFunction, error) {
	if root == "" {
		CurrentDirectory, _ := os.Getwd()
		root = CurrentDirectory
	}

	var functions []PHPFunction
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".php") {
			// Open the PHP file
			f, err := os.Open(path)
			if err != nil {
				fmt.Println("❌ ", err)
				return err
			}
			defer f.Close()
			// Use a regular expression to find function definitions
			scanner := bufio.NewScanner(f)
			re := regexp.MustCompile(`function\s+([a-zA-Z0-9_]+)`)
			for scanner.Scan() {
				line := scanner.Text()
				matches := re.FindStringSubmatch(line)
				if len(matches) > 1 {
					function := PHPFunction{
						Name: matches[1],
						Path: path,
					}
					functions = append(functions, function)
				}
			}
			if err := scanner.Err(); err != nil {
				fmt.Println("❌ ", err)
				return err
			}
		}
		return nil
	})
	return functions, err
}
