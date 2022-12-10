package library

import (
	"flag"
	"fmt"
)

var (
	GitFlag              = flag.Bool("git", false, "git mode")
	GitPushFlag          = flag.Bool("push", false, "commit and push")
	GitCommitMessageFlag = flag.String("m", "", "commit message")
)

/** Initiate Git Function */
func InitiateGitFunction() {
	/**
	 * Commit and Push
	 * - Equivalent to : `git commit -am "{message}" && git push origin HEAD`
	 */
	if *GitFlag && *GitPushFlag && *GitCommitMessageFlag != "" {
		fmt.Println("Commit and Push")
		cnp := fmt.Sprintf("git commit -am '%s'; git push origin HEAD", *GitCommitMessageFlag)
		cmd := [...]string{"bash", "-c", cnp}
		fmt.Println(ExecCommand(cmd[:]...))
	}
}
