package library

import (
	"flag"
	"fmt"
)

var (
	GitCommitandPushFlag = flag.Bool("git-cnp", false, "commit and push")
	GitCommitMessageFlag = flag.String("m", "", "commit message")
)

/** Initiate Git Function */
func InitiateGitFunction() {
	/** Commit and Push */
	if *GitCommitandPushFlag && *GitCommitMessageFlag != "" {
		fmt.Println("Commit and Push")
		find := fmt.Sprintf("git commit -am '%s'; git push origin HEAD", *GitCommitMessageFlag)
		cmd := [...]string{"bash", "-c", find}
		ExecCommand(cmd[:]...)
	}
}
