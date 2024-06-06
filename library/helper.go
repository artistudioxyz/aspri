package library

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

// Read File
func ReadFile(path string) []byte {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

// Rename File
func RenameFile(oldPath string, newPath string) {
	err := os.Rename(oldPath, newPath)
	if err != nil {
		fmt.Println("Error renaming file:", err)
		return
	}

	fmt.Println("✅ Success rename file", oldPath, "to", newPath)
}

// Write File
func WriteFile(FilePath string, content string) {
	f, err := os.Create(FilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err2 := f.WriteString(content)
	if err2 != nil {
		log.Fatal(err2)
	}
}

// Run custom bin command
func ExecCommand(args ...string) string {
	cmd := exec.Command(args[0], args[1:]...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Runing failed: %v", err)
	}
	return string(b)
}

// Slice Contains String
func SliceContainsString(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

// Get Shell Remove Function
func GetShellRemoveFunction(path string) string {
	if strings.Contains(path, "*") {
		PathArray := strings.Split(path, "/")
		path = strings.Replace(path, "/"+PathArray[len(PathArray)-1], "", 1)
		return fmt.Sprintf(`find %s -name "%s" -type f -delete;`, path, PathArray[len(PathArray)-1])
	} else {
		return fmt.Sprintf("rm -rf %s;", path)
	}
}

// Call an API endpoint with Method GET
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
