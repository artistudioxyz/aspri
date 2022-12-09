package library

import (
	"log"
	"os/exec"
)

/* Run custom bin command */
func ExecCommand(args ...string) string {
	cmd := exec.Command(args[0], args[1:]...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Runing failed: %v", err)
	}
	return string(b[:])
}
