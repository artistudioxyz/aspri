package library

import (
	"fmt"
	"os/exec"
	"time"
)

/** Function to calculate contributions */
// TODO: still hasn't implemented yet
func CalculateContributions(path string, contributors []string) {
	// Create Dir if Not Exists
	cmd := exec.Command("sh", "-c", fmt.Sprintf("mkdir -p %s/Insights/ && chmod -R 777 %s", path, path))
	cmd.Run()
	fmt.Printf("Execute: %s\n", cmd)

	// Output Stats
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && echo \"# 🐙 CONTRIBUTIONS : %s\" > %s/Insights/contributions.md", path, time.Now(), path))
	cmd.Run()
	fmt.Printf("Execute: %s\n", cmd)
	for _, contributor := range contributors {
		cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && aspri --file --count --text \"[[%s]]\" >> %s/Insights/contributions.md", path, contributor, path))
		cmd.Run()
		fmt.Printf("Execute: %s\n", cmd)
	}
}
