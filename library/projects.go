package library

import (
	"fmt"
	"strings"
)

/* Refactor Plugin */
func RefactorPlugin(path string, fromName string, toName string) {
	fmt.Println("Refactor Plugins")

	regular := fmt.Sprintf(
		"find %s -type f -name '*' -exec sed -i '' s/%s/%s/g {} +;",
		path, fromName, toName)
	uppercase := fmt.Sprintf(
		"find %s -type f -name '*' -exec sed -i '' s/%s/%s/g {} +;",
		path, strings.ToUpper(fromName), strings.ToUpper(toName))
	lowercase := fmt.Sprintf(
		"find %s -type f -name '*' -exec sed -i '' s/%s/%s/g {} +;",
		path, strings.ToLower(fromName), strings.ToLower(toName))

	cmd := [...]string{"bash", "-c", regular, uppercase, lowercase}
	ExecCommand(cmd[:]...)
}
