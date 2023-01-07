package library

import (
	"bufio"
	"fmt"
	"io/ioutil"
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
	/** Remove PHP Function */
	if *flags.PHP && *flags.RemoveFunction && *flags.Path != "" && len(*flags.FunctionName) > 0 {
		// TODO: Fix this function
		
		//Name := *flags.FunctionName
		//RemovePHPDocComment(*flags.Path, "send")
		//RemovePHPFunctionsByName(*flags.Path, "send")
		//removePHPDocBlocks(*flags.Path)
		//block, _ := getPHPDocBlock(*flags.Path, "send")
		//fmt.Println(block)

		RemovePHPDocBlocks(*flags.Path)
		//addDocBlock(*flags.Path, "Email", "send")
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

/** Remove PHP Doc Blocks */
func RemovePHPDocBlocks(dir string) error {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("❌ ", err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".php" {
			return nil
		}
		b, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println("❌ ", err)
			return err
		}
		r := regexp.MustCompile(`(?m)/\*{2,}[\s\S]*?\*/`)
		result := r.ReplaceAllString(string(b), "")
		if err = ioutil.WriteFile(path, []byte(result), 0); err != nil {
			fmt.Println("❌ ", err)
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println("❌ ", err)
		return err
	}
	return nil
}

func RemovePHPFunctionsByName(path string, functionName string) error {
	// Read the file
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Convert the file contents to a string
	s := string(b)

	// Use a regular expression to find the function definition
	re := regexp.MustCompile(`(public|private)\s+function\s+` + functionName + `\s*\(.*\)\s*\{[^\}]*\}`)
	s = re.ReplaceAllString(s, "")

	// Write the modified string back to the file
	err = ioutil.WriteFile(path, []byte(s), 0644)
	if err != nil {
		return err
	}

	return nil
}

/** Remove PHP Function */
func RemovePHPFunctionsByName2(path string, functionNames []string) error {
	// Open the PHP file
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("❌ ", err)
		return err
	}
	defer f.Close()
	// Use a regular expression to find function definitions
	scanner := bufio.NewScanner(f)
	var re *regexp.Regexp
	if len(functionNames) > 1 {
		reStr := `function\s+(`
		for i, name := range functionNames {
			if i > 0 {
				reStr += "|"
			}
			reStr += name
		}
		reStr += ")"
		re = regexp.MustCompile(reStr)
	} else if len(functionNames) == 1 {
		re = regexp.MustCompile(fmt.Sprintf(`function\s+%s`, functionNames[0]))
	} else {
		return fmt.Errorf("❌ no function names provided")
	}
	var lines []string
	inFunction := false
	inDocComment := false
	for scanner.Scan() {
		line := scanner.Text()
		if re.MatchString(line) {
			inFunction = true
		}
		if !inFunction && !inDocComment {
			lines = append(lines, line)
		}
		if inFunction && strings.Contains(line, "}") {
			inFunction = false
		}
		if strings.Contains(line, "/**") {
			inDocComment = true
		}
		if strings.Contains(line, "*/") {
			inDocComment = false
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("❌ ", err)
		return err
	}
	// Write the modified lines back to the file
	f, err = os.Create(path)
	if err != nil {
		fmt.Println("❌ ", err)
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
