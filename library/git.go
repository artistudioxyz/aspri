package library

import (
	"fmt"
)

/** Initiate Git Function */
func InitiateGitFunction(flags Flag) {
	/**
	 * Commit and Push
	 * - Equivalent to : `git commit -am "{message}" && git push origin HEAD`
	 */
	if *flags.Git && *flags.Message != "" {
		fmt.Println("📟 Commit and Push")
		cnp := fmt.Sprintf("git commit -am '%s'; git push origin HEAD", *flags.Message)
		cmd := [...]string{"bash", "-c", cnp}
		fmt.Println(ExecCommand(cmd[:]...))
	}
	/**
	* Re-stage
	* - Equivalent to : `git rm -rf cached . && git add .`
	 */
	if *flags.Git && *flags.Restage {
		fmt.Println("🏓 Re-staged")
		cmd := [...]string{"bash", "-c", "git rm -rf --cached . && git add ."}
		fmt.Println(ExecCommand(cmd[:]...))
	}
}
