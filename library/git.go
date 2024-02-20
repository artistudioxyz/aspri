package library

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// Initiate Git Function
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
	* Reset
	* - Equivalent to : `git reset --hard && git clean -df`
	 */
	if *flags.Git && *flags.ResetCache {
		fmt.Println("🔥 Reset")
		cmd := [...]string{"bash", "-c", "git reset --hard && git clean -df"}
		fmt.Println(ExecCommand(cmd[:]...))
	}
	/**
	* Reset Cache
	* - Equivalent to : `git rm -rf cached . && git add .`
	 */
	if *flags.Git && *flags.ResetCache {
		fmt.Println("🏓 Re-staged")
		cmd := [...]string{"bash", "-c", "git rm -rf --cached . && git add ."}
		fmt.Println(ExecCommand(cmd[:]...))
	}
	/** Git Gone */
	if *flags.Git && *flags.Gone {
		fmt.Println("🧽 Git Gone")
		GitGone()
	}
}

// Git Gone Implementation in Go
func GitGone() {
	// Run the "git fetch --all --prune" command
	fetchOut, err := exec.Command("git", "fetch", "--all", "--prune").Output()
	fmt.Println(string(fetchOut))
	if err != nil {
		fmt.Println("❌ Error fetching branches:", err)
		return
	}

	// Run the "git branch -vv" command
	branchCmd := exec.Command("git", "branch", "-vv")
	branches, err := branchCmd.CombinedOutput()
	if err != nil {
		fmt.Printf("❌ Error running 'git branch': %v\n", err)
		return
	}

	// Parse the output of "git branch -vv"
	scanner := bufio.NewScanner(bytes.NewReader(branches))
	var branchesToDelete []string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ": gone]") {
			branchesToDelete = append(branchesToDelete, strings.Fields(line)[0])
		}
	}

	// Delete branches with upstream tracking information "gone]"
	if len(branchesToDelete) > 0 {
		for _, branch := range branchesToDelete {
			_, err := exec.Command("git", "branch", "-D", branch).Output()
			if err != nil {
				fmt.Println("❌ Error deleting branch:", branch)
			} else {
				fmt.Println("🧹 Successfully delete branch:", branch)
			}
		}
	} else {
		fmt.Println("🙏 No branches with upstream tracking information 'gone]' found.")
	}
}
